<<<<<<< HEAD
// Calculates sum of item prices
=======
>>>>>>> testing-git
function sumPrices(items) {
    let total = 0;
    for (let i = 0; i < items.length; i++) {
        total += items[i].cost;
    }
    return total;
}
