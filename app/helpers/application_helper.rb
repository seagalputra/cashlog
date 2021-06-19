module ApplicationHelper
  def currency_idr_format(currency)
    number_to_currency currency, unit: 'Rp ', separator: ',', delimiter: '.', precision: 0
  end

  def generate_transaction_icon(title)
    title[0].upcase
  end
end
