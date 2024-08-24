import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import PageNumbers from "../../src/components/RecentTransactionTable/PageNumber";
import "@testing-library/jest-dom";

describe("PageNumbers", () => {
  const mockSetCurrentPage = jest.fn();

  beforeEach(() => {
    mockSetCurrentPage.mockClear();
  });

  it("renders correct number of page buttons", () => {
    render(
      <PageNumbers totalPages={10} currentPage={0} setCurrentPage={mockSetCurrentPage} />
    );

    const pageButtons = screen.getAllByRole("listitem");
    expect(pageButtons).toHaveLength(4); // 3 pages + ellipsis + last page
  });

  it("handles click on a page button", () => {
    render(
      <PageNumbers totalPages={5} currentPage={1} setCurrentPage={mockSetCurrentPage} />
    );

    const pageButton = screen.getByText("2");
    fireEvent.click(pageButton);
    expect(mockSetCurrentPage).toHaveBeenCalledWith(1);
  });

  it("renders ellipsis when pages are more than maxVisiblePages", () => {
    render(
      <PageNumbers totalPages={10} currentPage={1} setCurrentPage={mockSetCurrentPage} />
    );

    expect(screen.getByText("...")).toBeInTheDocument();
    expect(screen.getByText("10")).toBeInTheDocument();
  });

  it("does not render ellipsis when pages are less than maxVisiblePages", () => {
    render(
      <PageNumbers totalPages={3} currentPage={1} setCurrentPage={mockSetCurrentPage} />
    );

    expect(screen.queryByText("...")).not.toBeInTheDocument();
  });
});
