/**
 * Suppose we could access yesterday's stock prices as an array, where:
 *  - The indices are the time in minutes past trade opening time, which was 9:30am local time.
 *  - The values are the price in dollars of Apple stock at that time.
 * So if the stock cost $500 at 10:30am, stocks[60] = 500. Write an efficient function that takes stocks and
 * returns the best profit that could be made from 1 purchase and 1 sale of stock.
 * Notes: you must buy before you sell (cannot buy and sell at the same time)
 * Gotchas: it isn't sufficient to just take the difference between the highest and lowest prices (highest price may come before loweset price)
 *          if stock value goes down all day, then profit may be negative
 *          this is solvable in O(n) time and O(1) space
 */

public class HighestStockPrice {
  public int getMaxProfit(int[] stocks) {
    if (stocks.length < 2) {
      throw new IllegalArgumentException("Getting a profit requires at least 2 prices");
    }

    int minPrice = stocks[0];
    int maxProfit = stocks[1] - stocks[0];

    for (int i = 1; i < stocks.length; i++) {
      int currentPrice = stocks[i];
      int potentialPrice = currentPrice - minPrice;
      maxProfit = Math.max(maxProfit, potentialPrice);
      minPrice = Math.min(minPrice, currentPrice);
    }

    return maxProfit;
  }
}
