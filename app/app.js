window.addEventListener('load', function () {
    data.load();
    data.sync();
});

page = {};
page.goto = function (name) {
    var page = document.querySelector('[data-type="' + name + '"]');
    if (page) {
        document.querySelector('main').replaceChildren(page.content.cloneNode(true));
        history.pushState({}, "", `/` + name);
        data.sync();
    }
}

data = {};
data.root = {};
data.load = function () {
    var root = localStorage.getItem('data');
    if (root) {
        data.edit("", JSON.parse(root));
    } else {
        ajax.get("")
    }

    page.goto(window.location.pathname.slice(1));

}
data.save = function () {
    localStorage.setItem('data', JSON.stringify(data.root));
}
data.get = function (path) {
    let num = Number(path);
    if (!isNaN(num)) {
        return num;
    }

    let data = window.data.root;
    for (let field of path.split('.')) {
        if (!(field in data)) {
            //console.log(data);
            //console.error(`${field} not found in ${path}`);
            return null
        }
        data = data[field];
    }
    return data;
}
data.edit = function (path, value) {
    if (typeof value == 'object' && !Array.isArray(value) && value != null) {
        if (path != "") {
            path += ".";
        }
        for (let key in value) {
            data.edit(path + key, value[key]);
        }
        data.sync(path.slice(0, -1));
        return;
    }

    let head = data.root;
    let keys = path.split('.');
    for (let i = 0; i < keys.length - 1; i++) {
        let key = keys[i];
        if (key in head) {
            head = head[key];
        } else {
            head = head[key] = {};
        }
    }
    head[keys[keys.length - 1]] = value;

    data.sync(path);
    data.save();
}
data.push = function (path, value) {
    let slice = data.get(path);
    value = data.get(value);

    if (!slice) {
        data.edit(path, [value]);
    } else {
        slice.push(value);
    }
    data.sync(path);
    data.save();
}
data.pull = function (path, index) {
    let slice = data.get(path);
    if (!slice) {
        return;
    } else {
        slice.splice(index - 1, 1);
        if (slice.length == 0) {
            data.edit(path, null);
        }
    }
    data.sync(path);
    data.save();
}
data.feed = function (path) {
    let suffix = "";
    if (path) suffix = `="${path}"`;
    let feeds = document.querySelectorAll(`[data-feed${suffix}]`);
    for (let i = 0; i < feeds.length; i++) {
        let feed = feeds[i];
        if (!feed.template) {
            feed.template = feed.children[0];
        }

        let list = data.get(feed.dataset.feed);
        if (!list) {
            feed.length = 0;
            feed.innerHTML = "";
            continue;
        }


        if (feed.length == list.length) {
            return;
        }
        feed.innerHTML = "";

        for (let i = 0; i < list.length; i++) {
            let item = feed.dataset.feed + "." + i;

            let listitem = document.createElement("li");
            listitem.appendChild(feed.template.content.cloneNode(true));

            //attach attributes from template
            let attributes = feed.template.attributes;
            for (let i = 0; i < attributes.length; i++) {
                let attribute = attributes[i];
                listitem.setAttribute(attribute.name, attribute.value);
            }

            feed.appendChild(listitem);

            data.calc(listitem, {
                "..value": item,
                "..index": i + 1,
            });
        }
        feed.length = list.length;
    }
}
data.sync = function (path) {
    data.feed(path);

    let suffix = "";
    if (path) suffix = `="${path}"`;
    var elements = document.querySelectorAll(`[data-view${suffix}],[data-sync${suffix}],[data-edit${suffix}]`);

    for (let i = 0; i < elements.length; i++) {
        let element = elements[i];
        data.calc(element);
    }

    if (path) {
        suffix = `~="${path}"`;
    }

    elements = document.querySelectorAll(`[data-args${suffix}],[data-hook${suffix}]`);
    for (let i = 0; i < elements.length; i++) {
        data.calc(elements[i]);
    }
}
data.calc = function (element, scope) {
    if (scope) {
        let children = element.children;
        for (let i = 0; i < children.length; i++) {
            data.calc(children[i], scope);
        }
        let attributes = element.attributes;
        for (let i = 0; i < attributes.length; i++) {
            //replace with scope
            let attribute = attributes[i];
            let value = attribute.value;
            for (let key in scope) {
                value = value.replace(key, scope[key]);
            }
            element.setAttribute(attribute.name, value);
        }
    }

    let path = element.dataset.view || element.dataset.sync;
    if (path) {
        let value = data.get(element.dataset.view || element.dataset.sync);

        if (element.tagName == 'INPUT') {
            element.value = value;
        } else {
            element.innerHTML = value;
        }
    }

    if (element.dataset.edit || element.dataset.sync) {
        element.oninput = function () {
            data.edit(element.dataset.edit || element.dataset.sync, element.value);
        };
    }

    if (element.dataset.calc) {
        let fn = new Function(element.dataset.calc);

        let touching = {};

        fn.apply({
            attr: function (path, name, value) {

                let boolean = true;
                if (path.startsWith('!')) {
                    path = path.slice(1);
                    boolean = !boolean;
                }

                if (!("backup_" + name in element)) {
                    if (element.hasAttribute(name)) {
                        element["backup_" + name] = element.getAttribute(name);
                    } else {
                        element["backup_" + name] = null;
                    }
                }

                let condition = data.get(path);
                if (Array.isArray(condition)) {
                    condition = Boolean(condition.length > 0);
                }

                if (Boolean(condition) == boolean) {

                    if (value === true) {
                        element.setAttribute(name, "");
                    } else if (value === false) {
                        element.removeAttribute(name);
                    } else {
                        element.setAttribute(name, value);
                    }
                    touching[name] = true;

                } else {
                    if (!touching[name]) {
                        touching[name] = false;
                    }
                }
            },

            echo: function () {
                let args = arguments;
                if (!("format" in element)) {
                    element.format = element.innerText;
                }

                let text = element.format;
                for (let i = 0; i < args.length; i++) {
                    let arg = data.get(args[i]);
                    text = text.replace('{' + i + '}', arg);
                }
                element.innerText = text;
            },

        }, element);


        for (let attr in touching) {
            if (touching[attr]) {
                continue;
            }

            let backup = element["backup_" + attr];
            if (backup === null) {
                element.removeAttribute(attr);
            } else {
                element.setAttribute(attr, backup);
            }
        }
    }
}

ajax = {};
ajax.get = function (pointer) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', "/data/" + pointer, true);
    xhr.onload = function () {
        if (xhr.status == 200) {
            data.edit(pointer, JSON.parse(xhr.responseText));
            data.sync();
        }
    };
    xhr.send();
}
ajax.search = function (pointer) {
    var xhr = new XMLHttpRequest();
    xhr.open('SEARCH', "/data/" + pointer, true);
    xhr.onload = function () {
        if (xhr.status == 200) {
            data.edit(pointer, JSON.parse(xhr.responseText));
            data.sync();
        }
    };
    xhr.send(JSON.stringify(data.get(pointer)));
}
ajax.post = function (pointer) {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', "/data/" + pointer, true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function () {
        if (xhr.status == 200) {
            data.edit(pointer, JSON.parse(xhr.responseText));
            data.sync();
        }
    };
    xhr.send(JSON.stringify(data.get(pointer)));
};
ajax.put = function (pointer) {
    var xhr = new XMLHttpRequest();
    xhr.open('PUT', "/data/" + pointer, true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.send(JSON.stringify(data.get(pointer)));
};
ajax.delete = function (pointer) {
    var xhr = new XMLHttpRequest();
    xhr.open('DELETE', "/data/" + pointer, true);
    xhr.send();
}
ajax.patch = function (pointer, patch) {
    var xhr = new XMLHttpRequest();
    xhr.open('PATCH', "/data/" + pointer, true);
    xhr.setRequestHeader('Content-Type', 'application/json-patch+json; charset=UTF-8');
    xhr.send(JSON.stringify(patch));
}