import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import AllTransactionTable from "../../src/components/RecentTransactionTable/AllTransactionTable";
import { useGetAllTransactionsQuery } from "@/lib/redux/slices/transactionSlice";
import "@testing-library/jest-dom";

// Mock the useGetAllTransactionsQuery hook
jest.mock("../../src/lib/redux/slices/transactionSlice", () => ({
  useGetAllTransactionsQuery: jest.fn(),
}));

describe("AllTransactionTable", () => {
  const mockSetCurrentPage = jest.fn();

  beforeEach(() => {
    jest.clearAllMocks();
  });

  it("renders loading state correctly", () => {
    (useGetAllTransactionsQuery as jest.Mock).mockReturnValue({
      data: null,
      error: null,
      isLoading: true,
    });

    render(<AllTransactionTable />);
    expect(screen.getByText("Loading...")).toBeInTheDocument();
  });

  it("renders no transactions found state correctly", () => {
    (useGetAllTransactionsQuery as jest.Mock).mockReturnValue({
      data: { data: { content: [], totalPages: 0 } },
      error: null,
      isLoading: false,
    });

    render(<AllTransactionTable />);
    expect(screen.getByText("No transactions found.")).toBeInTheDocument();
  });

  it("renders transaction data correctly", () => {
    const mockData = {
      data: {
        content: [
          {
            amount: 100,
            description: "Test Transaction 1",
            transactionId: "123456789",
            type: "Income",
            date: "2023-01-01",
          },
        ],
        totalPages: 1,
      },
    };

    (useGetAllTransactionsQuery as jest.Mock).mockReturnValue({
      data: mockData,
      error: null,
      isLoading: false,
    });

    render(<AllTransactionTable />);

    expect(screen.getByText("Test Transaction 1")).toBeInTheDocument();
    expect(screen.getByText("123456789")).toBeInTheDocument();
    expect(screen.getByText("Income")).toBeInTheDocument();
    expect(screen.getByText("2023-01-01")).toBeInTheDocument();
  });

  it("renders Pagination component when totalPages > 1", () => {
    const mockData = {
      data: {
        content: [],
        totalPages: 2,
      },
    };

    (useGetAllTransactionsQuery as jest.Mock).mockReturnValue({
      data: mockData,
      error: null,
      isLoading: false,
    });

    render(<AllTransactionTable />);

    expect(screen.getByRole("navigation")).toBeInTheDocument();
  });
});
