Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html

  root 'dashboard#index'

  # get 'transactions', to: 'transactions#index', as: :transactions
  # get 'transactions/new', to: 'transactions#new', as: :transaction_new

  resources :transactions
end
