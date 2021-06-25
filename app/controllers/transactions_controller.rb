class TransactionsController < ApplicationController
  layout 'dashboard'

  def index
    @transactions = Transaction.all
  end

  def new; end

  def destroy
    @transaction = Transaction.find(params[:id])

    if @transaction.destroy
      redirect_to transactions_path, notice: 'Transaction successfully removed'
    else
      redirect_to transactions_path, alert: 'Failed to remove transaction'
    end
  end
end
