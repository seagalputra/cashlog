const toggleDropdown = () => {
  const userMenuButton = document.querySelector("#user-menu-button");
  userMenuButton.addEventListener("click", (event) => {
    event.preventDefault();

    const userMenu = document.querySelector("#user-menu");
    userMenu.classList.toggle("invisible");
  });
};

export { toggleDropdown };
