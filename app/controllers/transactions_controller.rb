class TransactionsController < ApplicationController

  def index
    @transactions = Transaction.all
  end

  def new
    @transaction = Transaction.new
  end

  def create
    redirect_to transactions_path, notice: 'Transaction successfully added'
  end

  def destroy
    @transaction = Transaction.find(params[:id])

    if @transaction.destroy
      redirect_to transactions_path, notice: 'Transaction successfully removed'
    else
      redirect_to transactions_path, alert: 'Failed to remove transaction'
    end
  end
end
