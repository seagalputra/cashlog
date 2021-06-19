class Transaction
  include Mongoid::Document
  include Mongoid::Timestamps

  validates :title, presence: true
  validates :price, presence: true

  field :title, type: String
  field :price, type: BigDecimal
  field :needs, type: BigDecimal, default: 0
  field :wants, type: BigDecimal, default: 0
  field :invest, type: BigDecimal, default: 0
  field :type, type: String
  field :notes, type: String

  def self.summaries
    [calculate_total(:needs), calculate_total(:wants), calculate_total(:invest)]
  end

  def self.calculate_total(type)
    all.pluck(type).map(&:to_i).reduce(0, :+)
  end
end
