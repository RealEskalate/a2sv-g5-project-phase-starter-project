import React from "react";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import Income from "../../src/components/RecentTransactionTable/Income";
import { useGetTransactionIncomeQuery } from "@/lib/redux/slices/transactionSlice";


// Mock the useGetTransactionIncomeQuery hook
jest.mock("../../src/lib/redux/slices/transactionSlice", () => ({
  useGetTransactionIncomeQuery: jest.fn(),
}));

describe("AllRecentIncomeTable", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it("renders loading state correctly", () => {
    (useGetTransactionIncomeQuery as jest.Mock).mockReturnValue({
      data: null,
      error: null,
      isLoading: true,
    });

    render(<Income />);
    expect(screen.getByText("Loading...")).toBeInTheDocument();
  });

  it("renders no transactions found state correctly", () => {
    (useGetTransactionIncomeQuery as jest.Mock).mockReturnValue({
      data: { data: { content: [], totalPages: 0 } },
      error: null,
      isLoading: false,
    });

    render(<Income />);
    expect(screen.getByText("No Income found.")).toBeInTheDocument();
  });

  it("renders transaction data correctly", () => {
    const mockData = {
      data: {
        content: [
          {
            amount: 100,
            description: "Test Income 1",
            transactionId: "123456789",
            type: "Deposit",
            date: "2023-01-01",
            card: "12345678"
          },
        ],
        totalPages: 1,
      },
    };

    (useGetTransactionIncomeQuery as jest.Mock).mockReturnValue({
        data:mockData,
        error: null,
        isLoading: false
    });

    render(<Income/>)

    expect(screen.getByText("Test Income 1")).toBeInTheDocument();
    expect(screen.getByText("+$100")).toBeInTheDocument();
    expect(screen.getByText("123456789")).toBeInTheDocument();
    expect(screen.getByText("Deposit")).toBeInTheDocument();
    expect(screen.getByText("2023-01-01")).toBeInTheDocument();
    expect(screen.getByText("1234 ****")).toBeInTheDocument();

  });

  it("renders Pagination component when totalPages > 1", () => {
    const mockData = {
      data: {
        content: [],
        totalPages: 2,
      },
    };

    (useGetTransactionIncomeQuery as jest.Mock).mockReturnValue({
      data: mockData,
      error: null,
      isLoading: false,
    });

    render(<Income />);
    expect(screen.getByRole("navigation")).toBeInTheDocument();
  });


});
