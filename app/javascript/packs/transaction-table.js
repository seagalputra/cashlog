function tableModal() {
	const actionButton = document.querySelectorAll("#table-action");

	actionButton.forEach((action, index) => {
		action.addEventListener("click", () => {
			const modal = document.querySelector(`#modal-${index}`);

			modal.classList.toggle("hidden");
		});
	});
}

export default tableModal;
