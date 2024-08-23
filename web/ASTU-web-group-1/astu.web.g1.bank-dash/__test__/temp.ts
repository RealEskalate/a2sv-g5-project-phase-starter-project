// import "@testing-library/jest-dom";
// import { render, screen } from "@testing-library/react";
// import MyCard from "@/components/MyCard/MyCard"; // Adjust import path if necessary

// describe("MyCard Component", () => {
//   it("should render card with correct static content", () => {
//     const mockProps = {
//       id: "1",
//       cardHolder: "Eddy Cusuma",
//       semiCardNumber: "37781234",
//       cardType: "Visa",
//       expiryDate: "2022-12-12",
//     };

//     render(<MyCard props={mockProps} index={0} />);

//     // Check if the text content is present
//     expect(screen.getByText('Balance')).toBeInTheDocument();
//     expect(screen.getByText('CARD HOLDER')).toBeInTheDocument();
//     expect(screen.getByText('Eddy Cusuma')).toBeInTheDocument();
//     expect(screen.getByText('VALID THRU')).toBeInTheDocument();
//     expect(screen.getByText('12/22')).toBeInTheDocument(); // Formatted expiry date
//   });

//   it("should render the chip card image", () => {
//     const mockProps = {
//       id: "1",
//       cardHolder: "Eddy Cusuma",
//       semiCardNumber: "37781234",
//       cardType: "Visa",
//       expiryDate: "2022-12-12",
//     };

//     render(<MyCard props={mockProps} index={0} />);

//     // Check if the image is present and has the correct src
//     const image = screen.getByAltText('chip_card');
//     expect(image).toBeInTheDocument();
//     expect(image).toHaveAttribute('src', '/assets/icons/chip-card-white.svg');
//   });

//   it("should render card with correct styles", () => {
//     const mockProps = {
//       id: "1",
//       cardHolder: "Eddy Cusuma",
//       semiCardNumber: "37781234",
//       cardType: "Visa",
//       expiryDate: "2022-12-12",
//     };

//     const { container } = render(<MyCard props={mockProps} index={0} />);

//     // Check if the styles are correctly applied (for example, background color)
//     expect(container.firstChild).toHaveClass('bg-grad-end'); // Assuming index 0 uses this class
//     expect(container.firstChild).toHaveClass('text-white');
//   });
// });
