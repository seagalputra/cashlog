class DashboardController < ApplicationController
  def index
    @needs, @wants, @invest = Transaction.summaries
    @balance = Transaction.calculate_total(:price)
    @recent_transactions = Transaction.all.limit(5)

    render inertia: 'Dashboard',
           props: { recent_transactions: @recent_transactions,
                    balance: @balance,
                    summary: { needs: @needs, wants: @wants, invest: @invest } }
  end
end
