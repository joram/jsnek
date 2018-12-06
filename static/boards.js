function get_square(x, y) {
    square_id = "square_" + x + "_" + y;
    square = $(`div#${square_id}`);
    return square;
}

function render_board(req) {
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
    for (key in b.data) {
        for (x in b.data[key]) {
            for (y in b.data[key][x]) {
                val = b.data[key][x][y];
                square = get_square(x, y);
                square.attr("data-dist", "d"+val);
            }
        }
    }
}