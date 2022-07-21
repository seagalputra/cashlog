Rails.application.routes.draw do
  # devise_for :users
  resources :expenses, except: [:show]

  root 'expenses#index'
end
