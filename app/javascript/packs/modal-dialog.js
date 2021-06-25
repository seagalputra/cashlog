import Rails from "@rails/ujs";
import Swal from "sweetalert2";

const nativeConfirm = Rails.confirm;

let __SkipConfirmation = false;

Rails.confirm = function (message, element) {
	if (__SkipConfirmation) return true;

	if (!element.hasAttribute("data-sweet-alert"))
		return nativeConfirm(message, element);

	function onConfirm() {
		__SkipConfirmation = true;
		element.click();
		__SkipConfirmation = false;
	}

	showDialog(element, onConfirm);

	return false;
};

function showDialog(element, onConfirm) {
	const options = JSON.parse(element.getAttribute("data-sweet-alert") || "{}");
	const message = element.getAttribute("data-confirm");

	Swal.fire({
		title: "Are you sure?",
		text: message,
		icon: "warning",
		showCancelButton: true,
		confirmButtonText: "Confirm",
		confirmButtonColor: "#3085d6",
		cancelButtonColor: "#d33",
		...options,
	}).then((result) => {
		if (result.isConfirmed) onConfirm();
	});
}
