page = {};
page.goto = function (name) {
    var page = document.querySelector('[data-type="' + name + '"]');
    if (page) {
        document.querySelector('main').replaceChildren(page.content.cloneNode(true));
        history.pushState({}, "", `/`+name);
        data.sync();
    }
}

title = function (s) {
    return s.charAt(0).toUpperCase() + s.slice(1);
}

data = {};
data.root = {};
data.load = function () {
    var root = localStorage.getItem('data');
    if (root) {
        data.set("", JSON.parse(root));
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
        if (!(field in data) && title(field) in data) {
            field = title(field);
        }
        if (!(field in data)) {
            console.log(data);
            console.error(`${field} not found in ${path}`);
            return null
        }
        data = data[field];
    }
    return data;
}
data.set = function (path, value) {
    if (path === "") {
        window.data.root = value;
        data.sync();
        data.save();
        return;
    }
    let split = path.split('.')
    if (split.length == 1) {
        data.root[split[0]] = value;
    } else {
        let base = split.slice(0, -1).join('.');
        data.get(base)[split[split.length - 1]] = value;
    }
    data.sync(path);
    data.save();
}
data.push = function (path, value) {
    let slice = data.get(path);
    if (!slice) {
        data.set(path, [value]);
    } else {
        slice.push(value);
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
            continue;
        }

        if (feed.length == list.length) {
            return; //TODO refresh
        }
        feed.innerHTML = "";

        for (let i = 0; i < list.length; i++) {
            let item = feed.dataset.feed+"."+i;

            for (let j = 0; j < feed.template.length; j++) {
                let element = feed.template[j].cloneNode(true);

                for (let name in element.dataset) {
                    switch (name) {
                        case 'feed':
                        case 'sync':
                        case 'view':
                        case 'args':
                        case 'scan':
                            element.dataset[name] = element.dataset[name].replace('..v', item)
                            element.dataset[name] = element.dataset[name].replace('..o', i)
                            element.dataset[name] = element.dataset[name].replace('..i', i+1)
                    }
    
                    let rename = name.replace('..v', item);
                    rename = rename.replace('..o', i);
                    rename = rename.replace('..i', i+1);
                    if (rename != name) {
                        element.dataset[rename] = element.dataset[name];
                        delete element.dataset[name];
                    }
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

        if (!element.dataset[backup] && !(backup + ":0" in element.dataset)) {
            if (element.hasAttribute(attr)) {
                element.dataset[backup] = element.getAttribute(attr);
            } else {
                element.dataset[backup + ":0"] = "";
            }
        }
        if (condition) {
            element.setAttribute(attr, value);
        } else {
            if (backup in element.dataset) {
                element.setAttribute(attr, element.dataset[backup]);
                delete element.dataset[backup];
            } else if (backup + ":0" in element.dataset) {
                element.removeAttribute(attr);
                delete element.dataset[backup + ":0"];
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

            if (args.length > 0) {
                if (args[0] == '0') {
                    condition = !data.get(path);
                } else {
                    condition = Boolean(data.get(path) == data.get(args[0]));
                }
            } else {
                condition = Boolean(data.get(path));
            }
            change(condition, element, attr, element.dataset[name]);
        }
    }
}

ajax = {};
ajax.get = function (pointer) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', "/data/" + pointer, true);
    xhr.onload = function () {
        if (xhr.status == 200) {
            data.set(pointer, JSON.parse(xhr.responseText));
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
            data.set(pointer, JSON.parse(xhr.responseText));
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
            data.set(pointer, JSON.parse(xhr.responseText));
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