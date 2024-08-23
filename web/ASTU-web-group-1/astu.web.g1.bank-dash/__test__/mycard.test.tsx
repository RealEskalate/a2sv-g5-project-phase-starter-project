import "@testing-library/jest-dom";
import { render, screen, fireEvent } from "@testing-library/react";
import MyCard from "@/components/MyCard/MyCard"; // Adjust import path if necessary

describe("MyCard Component", () => {
  it("should render card with correct static content", () => {
    render(<MyCard />);

    // Check if the text content is present
    expect(screen.getByText("Balance")).toBeInTheDocument();
    expect(screen.getByText("$5,756")).toBeInTheDocument();
    expect(screen.getByText("CARD HOLDER")).toBeInTheDocument();
    expect(screen.getByText("Eddy Cusuma")).toBeInTheDocument();
    expect(screen.getByText("VALID THRU")).toBeInTheDocument();
    expect(screen.getByText("12/22")).toBeInTheDocument();
    expect(screen.getByText("3778 **** **** 1234")).toBeInTheDocument();
  });

  it("should render the chip card image", () => {
    render(<MyCard />);

    // Check if the image is present and has the correct src
    const image = screen.getByAltText("chip_card");
    expect(image).toBeInTheDocument();
    expect(image).toHaveAttribute("src", "/assets/icons/chip-card-white.svg");
  });

  it("should render card with correct styles", () => {
    render(<MyCard />);
  });
});
