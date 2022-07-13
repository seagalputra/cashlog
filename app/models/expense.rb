class Expense < ApplicationRecord
  OUTCOME = 'outcome'.freeze
  INCOME = 'income'.freeze

  enum expense_type: [:income, :outcome]
end
