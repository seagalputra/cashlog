class Transaction
  include Mongoid::Document
  include Mongoid::Timestamps

  field :title, type: String
  field :price, type: BigDecimal
  field :needs, type: BigDecimal
  field :wants, type: BigDecimal
  field :invest, type: BigDecimal
  field :notes, type: String
end
