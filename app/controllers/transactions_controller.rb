class TransactionsController < ApplicationController
  def index
    @transactions = Transaction.all
  end

  def new; end
end
