class DashboardController < ApplicationController
  def index
    @needs, @wants, @invest = Transaction.summaries
    @recent_transactions = Transaction.all.limit(5)
  end
end
