class Expense < ApplicationRecord
  TYPES = [:income, :outcome].freeze

  def income?
    expense_type.downcase == 'income'
  end
end
