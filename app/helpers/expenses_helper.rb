module ExpensesHelper
  def number_to_currency_idr(amount)
    number_to_currency(amount, unit: 'Rp', separator: ',', delimiter: '.', precision: 0)
  end
end
