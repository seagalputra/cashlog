Rails.application.routes.draw do
  resources :expenses
  # devise_for :users
  root 'home#index'
end
