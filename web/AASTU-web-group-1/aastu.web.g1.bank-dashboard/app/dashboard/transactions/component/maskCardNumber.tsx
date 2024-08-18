function maskCardNumber(cardNumber: string) {
  // Ensure the card number is a string
  const strCardNumber = cardNumber.toString();

  const length = strCardNumber.length;

  if (length <= 8) {
    const firstFour = strCardNumber.slice(0, 4);
    const lastFour = strCardNumber.slice(-4);
    const maskedSection = "****";
    return `${firstFour} ${maskedSection} ${lastFour}`;
  }

  const firstFour = strCardNumber.slice(0, 4);
  const lastFour = strCardNumber.slice(-4);
  const maskedSection = strCardNumber
    .slice(4, -4)
    .replace(/./g, "*")
    .replace(/(.{4})/g, "$1 ");

  return `${firstFour} ${maskedSection.trim()} ${lastFour}`;
}
 
export default maskCardNumber;