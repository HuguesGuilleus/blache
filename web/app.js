"use strict";
var __values = (this && this.__values) || function(o) {
    var s = typeof Symbol === "function" && Symbol.iterator, m = s && o[s], i = 0;
    if (m) return m.call(o);
    if (o && typeof o.length === "number") return {
        next: function () {
            if (o && i >= o.length) o = void 0;
            return { value: o && o[i++], done: !o };
        }
    };
    throw new TypeError(s ? "Object is not iterable." : "Symbol.iterator is not defined.");
};
var __read = (this && this.__read) || function (o, n) {
    var m = typeof Symbol === "function" && o[Symbol.iterator];
    if (!m) return o;
    var i = m.call(o), r, ar = [], e;
    try {
        while ((n === void 0 || n-- > 0) && !(r = i.next()).done) ar.push(r.value);
    }
    catch (error) { e = { error: error }; }
    finally {
        try {
            if (r && !r.done && (m = i["return"])) m.call(i);
        }
        finally { if (e) throw e.error; }
    }
    return ar;
};
var Coordinate = {
    parse: function (s) {
        var t = s.split(',');
        if (t.length != 2)
            throw "Invalid syntax for coordinates string";
        return [parseInt(t[0]), parseInt(t[1])];
    },
    toString: function (x, z) {
        return x + ',' + z;
    },
};
// Download d for the user in the file name.
function userDownload(b, name) {
    var u = URL.createObjectURL(b);
    var a = document.createElement('a');
    document.body.appendChild(a);
    a.href = u;
    a.download = name;
    a.click();
    a.remove();
    URL.revokeObjectURL(u);
}
var REGION_SIZE = 16 * 32;
var blackImage = function () {
    var raw = new Uint8ClampedArray(REGION_SIZE * REGION_SIZE * 4);
    for (var i = 0; i < raw.length; i += 4) {
        raw[i + 3] = 255;
    }
    return new ImageData(raw, REGION_SIZE, REGION_SIZE);
}();
var TileType;
(function (TileType) {
    TileType["bloc"] = "bloc";
    TileType["biome"] = "biome";
    TileType["height"] = "height";
})(TileType || (TileType = {}));
var Viewer = /** @class */ (function () {
    function Viewer() {
        var _a;
        var _this = this;
        this.regions = (_a = {},
            _a[TileType.bloc] = new Map(),
            _a[TileType.biome] = new Map(),
            _a[TileType.height] = new Map(),
            _a);
        this.struct = new Map();
        this.type = TileType.bloc;
        this.posX = 0;
        this.posZ = 0;
        this.size = REGION_SIZE;
        // Use to get some element into the dom.
        function $(id) {
            var e = document.getElementById(id);
            if (!e)
                throw "Element '".concat(id, "' no found in DOM");
            return e;
        }
        // Elements
        this.coordsRegion = $('coordsRegion');
        this.coordsBloc = $('coordsBloc');
        this.coordsStruct = $('coordsStruct');
        this.canvas_el = $('canvas2d');
        var c = this.canvas_el.getContext('2d');
        if (!c)
            throw 'Can not create canvas';
        this.canvas_ctx = c;
        $('tileTypeBloc')
            .addEventListener('click', function () { return _this.typeChange(TileType.bloc); });
        $('tileTypeBiome')
            .addEventListener('click', function () { return _this.typeChange(TileType.biome); });
        $('tileTypeHeight')
            .addEventListener('click', function () { return _this.typeChange(TileType.height); });
        var frontier = $('enableFrontier');
        this.enableFrontier = frontier.checked;
        frontier.addEventListener('change', function () {
            _this.enableFrontier = frontier.checked;
            _this.drawAll();
        });
        var structs = $('enableStruct');
        this.enableStruct = structs.checked;
        structs.addEventListener('change', function () {
            _this.enableStruct = structs.checked;
            _this.drawAll();
        });
        // Download regions list
        fetch('regions.json')
            .then(function (rep) { return rep.json(); })
            .then(function (rs) {
            var e_1, _a;
            try {
                for (var rs_1 = __values(rs), rs_1_1 = rs_1.next(); !rs_1_1.done; rs_1_1 = rs_1.next()) {
                    var c_1 = rs_1_1.value;
                    _this.regions[TileType.bloc].set(c_1, null);
                    _this.regions[TileType.biome].set(c_1, null);
                    _this.regions[TileType.height].set(c_1, null);
                }
            }
            catch (e_1_1) { e_1 = { error: e_1_1 }; }
            finally {
                try {
                    if (rs_1_1 && !rs_1_1.done && (_a = rs_1.return)) _a.call(rs_1);
                }
                finally { if (e_1) throw e_1.error; }
            }
            _this.drawAll();
        });
        window.addEventListener('load', function () { return _this.resize(); });
        window.addEventListener('resize', function () { return _this.resize(); });
        window.addEventListener('keydown', function (e) {
            switch (e.key) {
                case '-':
                    _this.zoomOut(_this.canvas_el.width / 2, _this.canvas_el.height / 2);
                    break;
                case '+':
                    _this.zoomIn(_this.canvas_el.width / 2, _this.canvas_el.height / 2);
                    break;
                case 'ArrowLeft':
                    _this.posX -= _this.size / 4;
                    break;
                case 'ArrowRight':
                    _this.posX += _this.size / 4;
                    break;
                case 'ArrowUp':
                    _this.posZ -= _this.size / 4;
                    break;
                case 'ArrowDown':
                    _this.posZ += _this.size / 4;
                    break;
                case '0':
                    _this.posX = _this.posZ = 0;
                    _this.size = REGION_SIZE;
                    break;
                case 's':
                    _this.canvas_el.toBlob(function (b) { return userDownload(b !== null && b !== void 0 ? b : new Blob(), "".concat(document.location.hostname, "_").concat(Math.trunc(REGION_SIZE / _this.size), "_").concat(Math.floor(_this.posX / _this.size), ".").concat(Math.floor(_this.posZ / _this.size), ".png")); });
                    return;
                default:
                    return;
            }
            _this.coordsBloc.innerText = '?';
            _this.coordsRegion.innerText = '?';
            _this.drawAll();
        });
        this.canvas_el.addEventListener('mousemove', function (event) {
            var e_2, _a;
            var _b;
            var x = Math.trunc((_this.posX + event.x) * REGION_SIZE / _this.size), z = Math.trunc((_this.posZ + event.y) * REGION_SIZE / _this.size), rx = Math.floor(x / REGION_SIZE), rz = Math.floor(z / REGION_SIZE), px = x - rx * REGION_SIZE, pz = z - rz * REGION_SIZE;
            _this.coordsBloc.innerText = "".concat(x, ", ").concat(z);
            _this.coordsRegion.innerText = "".concat(rx, ", ").concat(rz);
            _this.coordsStruct.innerText = '';
            try {
                for (var _c = __values(((_b = _this.struct.get(Coordinate.toString(rx, rz))) !== null && _b !== void 0 ? _b : [])), _d = _c.next(); !_d.done; _d = _c.next()) {
                    var s = _d.value;
                    if (Math.abs(px - s.x * 16) < 10 && Math.abs(pz - s.z * 16) < 10) {
                        _this.coordsStruct.innerText = s.name;
                        break;
                    }
                }
            }
            catch (e_2_1) { e_2 = { error: e_2_1 }; }
            finally {
                try {
                    if (_d && !_d.done && (_a = _c.return)) _a.call(_c);
                }
                finally { if (e_2) throw e_2.error; }
            }
            if (!event.buttons)
                return;
            _this.posX -= event.movementX;
            _this.posZ -= event.movementY;
            _this.drawAll();
        });
        this.canvas_el.addEventListener('wheel', function (event) {
            var d = event.deltaY;
            if (d > 0) {
                _this.zoomOut(event.x, event.y);
            }
            else if (d < 0) {
                _this.zoomIn(event.x, event.y);
            }
            _this.drawAll();
        });
    }
    // Edit the type of tile
    Viewer.prototype.typeChange = function (t) {
        if (t != this.type) {
            this.type = t;
            this.drawAll();
        }
    };
    Viewer.prototype.zoomIn = function (w, h) {
        if (this.size > REGION_SIZE * 16)
            return;
        this.posX = this.posX * 2 + w;
        this.posZ = this.posZ * 2 + h;
        this.size *= 2;
    };
    Viewer.prototype.zoomOut = function (w, h) {
        if (this.size < REGION_SIZE / 16)
            return;
        this.posX = this.posX / 2 - w / 2;
        this.posZ = this.posZ / 2 - h / 2;
        this.size /= 2;
    };
    // Change the size of the canvas.
    Viewer.prototype.resize = function () {
        this.canvas_el.width = window.innerWidth;
        this.canvas_el.height = window.innerHeight;
        this.drawAll();
    };
    // Draw all regions.
    Viewer.prototype.drawAll = function () {
        var e_3, _a;
        this.canvas_ctx.fillStyle = 'black';
        this.canvas_ctx.fillRect(0, 0, this.canvas_el.width, this.canvas_el.height);
        var X = this.posX;
        var Z = this.posZ;
        var S = this.size;
        var W = this.canvas_el.width;
        var H = this.canvas_el.height;
        var T = this.type;
        try {
            // Draw regions
            for (var _b = __values(this.regions[this.type].keys()), _c = _b.next(); !_c.done; _c = _b.next()) {
                var k = _c.value;
                var _d = __read(Coordinate.parse(k), 2), x = _d[0], z = _d[1];
                if ((x + 1) * S > X
                    && x * S < X + W
                    && (z + 1) * S > Z
                    && z * S < Z + H) {
                    this.getTile(k, x, z, T);
                }
            }
        }
        catch (e_3_1) { e_3 = { error: e_3_1 }; }
        finally {
            try {
                if (_c && !_c.done && (_a = _b.return)) _a.call(_b);
            }
            finally { if (e_3) throw e_3.error; }
        }
    };
    // Get the region tile and the region's structures, then draw all this info.
    Viewer.prototype.getTile = function (k, x, z, t) {
        this.drawTile(x, z, t, this.getRegion(k, x, z, t), this.getStructures(k, x, z));
    };
    // Return an image for the region. If the region is'nt yet fetch, fetch it
    // and call View.getTile() when the image is loaded.
    Viewer.prototype.getRegion = function (k, x, z, t) {
        var _this = this;
        var img = this.regions[t].get(k);
        if (img != null)
            return img;
        this.regions[t].set(k, blackImage);
        var i = new Image(REGION_SIZE, REGION_SIZE);
        i.onload = function () {
            _this.regions[t].set(k, i);
            _this.getTile(k, x, z, t);
        };
        i.src = "".concat(t, "/").concat(x, ".").concat(z, ".png");
        return blackImage;
    };
    // Return the list of structure, if the list is not yet fetch, fetch it
    // then call View.getTile().
    Viewer.prototype.getStructures = function (k, x, z) {
        var _this = this;
        var l = this.struct.get(k);
        if (Array.isArray(l))
            return l;
        this.struct.set(k, []);
        fetch("structs/".concat(x, ".").concat(z, ".json"))
            .then(function (rep) { return rep.json(); })
            .then(function (l) {
            _this.struct.set(k, l.map(function (s) {
                s.color = "hsl(\n\t\t\t\t\t\t".concat(s.name.split('').reduce(function (s, c) { return (s + c.charCodeAt(0)) % 256; }, 0), ", 100%, 50%)");
                return s;
            }));
            _this.getTile(k, x, z, _this.type);
        });
        return [];
    };
    // Draw the region image into the screen.
    Viewer.prototype.drawTile = function (xa, za, t, surface, struct) {
        var e_4, _a;
        if (t != this.type)
            return;
        var S = this.size, xr = xa * S - this.posX, zr = za * S - this.posZ;
        if (surface instanceof HTMLImageElement) {
            this.canvas_ctx.imageSmoothingEnabled = false;
            this.canvas_ctx.drawImage(surface, xr, zr, S, S);
        }
        else {
            this.canvas_ctx.putImageData(surface, xr, zr, 0, 0, S, S);
        }
        if (this.enableFrontier && S > 32) {
            if (S > 512) {
                this.canvas_ctx.strokeStyle = 'red';
                this.canvas_ctx.strokeStyle = 'yellow';
                this.canvas_ctx.strokeStyle = 'grey';
                this.canvas_ctx.lineWidth = 1;
                for (var i = 1; i < 32; i++) {
                    this.canvas_ctx.stroke(new Path2D("M".concat(xr + S / 32 * i, " ").concat(zr, " v").concat(S)));
                    this.canvas_ctx.stroke(new Path2D("M".concat(xr, " ").concat(zr + S / 32 * i, " h").concat(S)));
                }
            }
            this.canvas_ctx.strokeStyle = 'orange';
            this.canvas_ctx.lineWidth = 2.5;
            this.canvas_ctx.stroke(new Path2D("M".concat(xr, " ").concat(zr, " v").concat(S, " h").concat(S, " v").concat(-S, " z")));
        }
        if (this.enableStruct) {
            try {
                for (var struct_1 = __values(struct), struct_1_1 = struct_1.next(); !struct_1_1.done; struct_1_1 = struct_1.next()) {
                    var s = struct_1_1.value;
                    this.canvas_ctx.fillStyle = s.color;
                    this.canvas_ctx.fillRect(xr + s.x * 16 * S / REGION_SIZE, zr + s.z * 16 * S / REGION_SIZE, 8, 8);
                }
            }
            catch (e_4_1) { e_4 = { error: e_4_1 }; }
            finally {
                try {
                    if (struct_1_1 && !struct_1_1.done && (_a = struct_1.return)) _a.call(struct_1);
                }
                finally { if (e_4) throw e_4.error; }
            }
        }
    };
    return Viewer;
}());
(function () {
    function main() {
        new Viewer();
    }
    document.readyState == 'loading' ? document.addEventListener('DOMContentLoaded', main, {
        once: true
    }) : main();
}());
