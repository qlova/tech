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
            console.log(data);
            console.error(`${field} not found in ${path}`);
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
            data.edit(path + key.toLowerCase(), value[key]);
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
            feed.template = [];
            while (feed.hasChildNodes()) {
                feed.template.push(feed.removeChild(feed.firstChild));
            }
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

            for (let j = 0; j < feed.template.length; j++) {
                let element = feed.template[j].cloneNode(true);

                for (let name in element.dataset) {
                    switch (name) {
                        case 'feed':
                        case 'sync':
                        case 'view':
                        case 'args':
                        case 'scan':
                            element.dataset[name] = element.dataset[name].replace('..value', item)
                            element.dataset[name] = element.dataset[name].replace('..index', i + 1)
                    }

                    let rename = name.replace('..value', item);
                    rename = rename.replace('..index', i + 1);
                    if (rename != name) {
                        element.dataset[rename] = element.dataset[name];
                        delete element.dataset[name];
                    }

                    element.index = i + 1;
                    element.value = item;
                }

                data.calc(element);
                feed.appendChild(element);
            }
        }

        feed.length = list.length;
    }
}
data.sync = function (path) {
    data.feed(path);

    let suffix = "";
    if (path) suffix = `="${path}"`;
    var elements = document.querySelectorAll(`[data-view${suffix}],[data-sync${suffix}]`);

    for (let i = 0; i < elements.length; i++) {
        let element = elements[i];
        data.calc(element);
    }

    if (path) {
        suffix = `~="${path}"`;
    }

    elements = document.querySelectorAll(`[data-args${suffix}]`);
    for (let i = 0; i < elements.length; i++) {
        data.calc(elements[i]);
    }
}
data.calc = function (element) {
    let path = element.dataset.view || element.dataset.sync;
    if (path) {
        let value = data.get(element.dataset.view || element.dataset.sync);

        if (element.tagName == 'INPUT') {
            element.value = value;
        } else {
            element.innerHTML = value;
        }
    }

    let change = function (condition, element, attr, value) {
        let backup = "backup:" + attr;

        if (!(backup in element.dataset) && !(backup + ":0" in element.dataset)) {
            if (element.hasAttribute(attr)) {
                element.dataset[backup] = element.getAttribute(attr);
            } else {
                element.dataset[backup + ":0"] = "";
            }
        }

        if (condition) {
            if (value == null) {
                element.removeAttribute(attr);
            } else {
                element.setAttribute(attr, value);
            }
        }
    }

    if ("echo" in element.dataset) {
        let args = element.dataset.args.split(' ');
        let text = element.dataset.echo;

        for (let i = 0; i < args.length; i++) {
            let arg = data.get(args[i]);
            text = text.replace('%' + i, arg);
        }

        element.innerText = text;
    }

    let touching = {};
    for (let name in element.dataset) {
        if (name.startsWith('when:')) {
            let split = name.substring(5).split(':');
            if (split.length < 2) {
                continue;
            }
            let path = split[0];
            let args = split.slice(1, split.length - 1);
            let attr = split[split.length - 1];
            let condition = false;
            let value = element.dataset[name];

            if (attr.charAt(0).toLowerCase() !== attr.charAt(0)) {
                attr = attr.toLowerCase();
                value = null;
            }
            

            if (args.length > 0) {
                if (args[0] == '0') {
                    condition = !data.get(path);
                } else {
                    condition = Boolean(data.get(path) == data.get(args[0]));
                }
            } else {
                condition = Boolean(data.get(path));
            }

            //console.log(condition, path, args, attr, value);
            change(condition, element, attr, value);
            if (condition) {
                touching[attr] = true;
            } else if (!(attr in touching)) {
                touching[attr] = false;
            }
        }
    }
    for (let attr in touching) {
        if (touching[attr]) {
            continue;
        }
        let backup = "backup:" + attr;

        if (backup in element.dataset) {
            element.setAttribute(attr, element.getAttribute["data-"+backup]);
            delete element.dataset[backup];
        } else if (backup + ":0" in element.dataset) {
            element.removeAttribute(attr);
            delete element.dataset[backup + ":0"];
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