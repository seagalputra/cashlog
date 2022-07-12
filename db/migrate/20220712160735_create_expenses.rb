class CreateExpenses < ActiveRecord::Migration[7.0]
  def change
    create_table :expenses do |t|
      t.string :title
      t.decimal :price
      t.date :date
      t.string :expense_type
      t.text :description

      t.timestamps
    end
  end
end
