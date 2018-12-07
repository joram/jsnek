function get_square(x, y) {
    square_id = "square_" + x + "_" + y;
    square = $(`div#${square_id}`);
    return square;
}

function render_slider(board_i, requests){
    if(board_i<0){
        board_i = 0;
    }
    if(board_i >= requests.length){
        board_i = requests.length-1
    }

    req = requests[board_i];
    b = req.board;
    board = $("#board");

    // frame slider
    s = $(`<div class="slider" style="grid-column-start:1; grid-column-end:${b.width};"><input type="range" min="1" max="${req.length}" value="${board_i}" class="slider" id="frame_picker"></div>`);
    board.append(s);
    picker = $("#frame_picker");
    picker.on("change", function(){
        i = parseInt(picker.val());
        console.log("changing to frame "+i);
        render_board(requests[i]);
    });

    $(document).keydown(function(e) {
        switch(e.which) {
            case 37: // left
                render_board(board_i-1, requests);
                break;

            case 39: // right
                render_board(board_i+1, requests);
                break;
            default: return; // exit this handler for other keys
        }
        e.preventDefault(); // prevent the default action (scroll / move caret)
    });
}

function render_board(board_i, requests) {
    req = requests[board_i];
    b = req.board;
    board = $("#board");
    board.empty();

    // draw grid
    for (y = 0; y < b.height; y++) {
        for (x = 0; x < b.width; x++) {
            square_id = "square_" + x + "_" + y;
            c = x + 1;
            r = y + 1;
            square = $(`<div id='${square_id}' class='square' style="grid-column:${c}; grid-row:${r};"></div>`);
            board.append(square);
        }
    }

    // draw food
    for (i = 0; i < b.food.length; i++) {
        food = b.food[i];
        square = get_square(food.x, food.y);
        square.append($(`<div class='food'/>`));
    }

    // draw snakes
    for (i = 0; i < b.snakes.length; i++) {
        snake = b.snakes[i];
        for (j = 0; j < snake.body.length; j++) {
            coord = snake.body[j];
            square = get_square(coord.x, coord.y);
            cls = "snake ";
            if (j === 0) {
                cls += "head "
            }
            if (snake.id ===
                req.you.id) {
                cls += " my"
            }
            square.append($(`<div class='${cls}'/>`));
        }
    }

    // draw snake distances
    key = "time_to_"+;
    for (x in b.data[key]) {
        for (y in b.data[key][x]) {
            val = b.data[key][x][y];
            square = get_square(x, y);
            square.attr("data-dist", "d"+val);
        }
    }


    render_slider(board_i, requests);
}