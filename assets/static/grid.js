/* globals Config */

/* exported Grid */
var Grid = (function() {

    var selectors = {
        alerts: "#alerts",
        incident: ".incident",
        gridSizer: ".grid-sizer",
    };

    var grid;

    var init = function() {
        grid = $(selectors.alerts).masonry({
            itemSelector: selectors.incident,
            columnWidth: selectors.gridSizer,
            percentPosition: true,
            transitionDuration: "0.4s",
            hiddenStyle: {
                opacity: 0
            },
            visibleStyle: {
                opacity: 1
            }
        });
    };

    var clear = function() {
        grid.masonry("remove", $(selectors.incident));
    };

    var redraw = function() {
        grid.masonry("layout");
    };

    var remove = function(elem) {
        grid.masonry("remove", elem);
    };

    var append = function(elem) {
        if (Config.GetOption("appendtop").Get()) {
            grid.prepend(elem).masonry("prepended", elem);
        } else {
            grid.append(elem).masonry("appended", elem);
        }
    };

    var items = function() {
        return grid.masonry("getItemElements");
    };

    var hide = function() {
        $(selectors.alerts).hide();
    };

    var show = function() {
        $(selectors.alerts).show();
    };

    return {
        Init: init,
        Clear: clear,
        Hide: hide,
        Show: show,
        Redraw: redraw,
        Append: append,
        Remove: remove,
        Items: items
    };

})();
