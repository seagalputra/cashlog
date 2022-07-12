json.extract! expense, :id, :title, :price, :date, :expense_type, :description, :created_at, :updated_at
json.url expense_url(expense, format: :json)
