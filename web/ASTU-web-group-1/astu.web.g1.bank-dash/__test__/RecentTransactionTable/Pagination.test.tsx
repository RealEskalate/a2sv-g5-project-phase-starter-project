import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import Pagination from "../../src/components/RecentTransactionTable/Pagination";
import "@testing-library/jest-dom";

describe("Pagination", () => {
  const mockSetCurrentPage = jest.fn();

  beforeEach(() => {
    mockSetCurrentPage.mockClear();
  });

  it("renders previous and next buttons", () => {
    render(
      <Pagination totalPages={5} currentPage={2} setCurrentPage={mockSetCurrentPage} />
    );

    expect(screen.getByText("Previous")).toBeInTheDocument();
    expect(screen.getByText("Next")).toBeInTheDocument();
  });

  it("previous button calls setCurrentPage with the previous page", () => {
    render(
      <Pagination totalPages={5} currentPage={2} setCurrentPage={mockSetCurrentPage} />
    );

    const prevButton = screen.getByText("Previous");
    fireEvent.click(prevButton);
    expect(mockSetCurrentPage).toHaveBeenCalledWith(1);
  });

  it("next button calls setCurrentPage with the next page", () => {
    render(
      <Pagination totalPages={5} currentPage={2} setCurrentPage={mockSetCurrentPage} />
    );

    const nextButton = screen.getByText("Next");
    fireEvent.click(nextButton);
    expect(mockSetCurrentPage).toHaveBeenCalledWith(3);
  });

  it("previous button is disabled on first page", () => {
    render(
      <Pagination totalPages={5} currentPage={1} setCurrentPage={mockSetCurrentPage} />
    );

    const prevButton = screen.getByText("Previous");
    expect(prevButton).toBeDisabled();
  });

  it("next button is disabled on last page", () => {
    render(
      <Pagination totalPages={5} currentPage={5} setCurrentPage={mockSetCurrentPage} />
    );

    const nextButton = screen.getByText("Next");
    expect(nextButton).toBeDisabled();
  });
});
