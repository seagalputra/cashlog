class DashboardController < ApplicationController
  def index
    @needs, @wants, @invest = Transaction.summaries
    @recent_transactions = Transaction.all.limit(5)

    render inertia: 'Dashboard',
           props: { recent_transactions: @recent_transactions,
                    summary: { needs: @needs, wants: @wants, invest: @invest } }
  end
end
