# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: 'Star Wars' }, { name: 'Lord of the Rings' }])
#   Character.create(name: 'Luke', movie: movies.first)

if Rails.env.development?
  puts 'Seeding transaction data into `transactions` document'

  25.times do
    type = %w[income outcome].sample
    category = %i[needs wants invest].sample
    price = type == 'income' ? rand(1_000..5_000_000) : rand(1_000..5_000_000) * -1

    initial_parition = { needs: 0, wants: 0, invest: 0 }
    partition = if type == 'income'
                  { needs: price * 0.5, wants: price * 0.3, invest: price * 0.2 }
                else
                  { **initial_parition, **Hash[category, price] }
                end

    Transaction.create({ **partition, title: Faker::Lorem.sentence(word_count: 2),
                                      price: price,
                                      type: type,
                                      notes: Faker::Lorem.paragraph(sentence_count: 2) })
  end
end
