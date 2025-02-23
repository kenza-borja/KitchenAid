//displays day and date
function updateDate() {
    const today = new Date();
    const options = { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' };
    const formattedDate = today.toLocaleDateString('en-GB', options).toUpperCase();

    document.getElementById("dateDisplay").textContent = formattedDate;
}

window.onload = updateDate;

//add item button changes the page
document.getElementById("add_item_button").addEventListener("click", function() {
    window.location.href = "add_item.html";
});
