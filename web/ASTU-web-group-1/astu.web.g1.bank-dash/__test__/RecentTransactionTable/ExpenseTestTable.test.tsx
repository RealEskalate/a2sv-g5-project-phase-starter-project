import React from "react";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import Expense from "../../src/components/RecentTransactionTable/Expense";
import { useGetTransactionExpenseQuery } from "@/lib/redux/slices/transactionSlice";

//Mock the useGetTransactionExpenseQuery hook
jest.mock("../../src/lib/redux/slices/transactionSlice", () => ({
  useGetTransactionExpenseQuery: jest.fn(),
}));

describe("AllRecentExpenseTable",() => {
    beforeEach(() => {
        jest.clearAllMocks();
    })

    it("renders loading state correctly", () => {
        (useGetTransactionExpenseQuery as jest.Mock).mockReturnValue({
            data: null,
            error: null,
            isLoading: true,
        });

        render(<Expense />);
        expect(screen.getByText("Loading...")).toBeInTheDocument();
    });

    it("renders error state correctly", () => {
        (useGetTransactionExpenseQuery as jest.Mock).mockReturnValue({
            data: null,
            error: { message: "Error fetching data" },
            isLoading: false,
        });

        render(<Expense />);
        expect(screen.getByText("Opps Something happens")).toBeInTheDocument();
    });

    it("renders transaction data correctly", () => {
        const mockData = {
            data : {

                content: [
                    {
                        amount: -100,
                        description: "Test Expense 1",
                        transactionId: "123456789",
                        type: "Expense",
                        date: "2023-01-01",
                        card: "12345678"
                    },
                ],
                totalPages: 1
            }
        };
        (useGetTransactionExpenseQuery as jest.Mock).mockReturnValue({
            data: mockData,
            error: null,
            isLoading: false,
        });

        render(<Expense />);
        expect(screen.getByText("Test Expense 1")).toBeInTheDocument();
        expect(screen.getByText("-$100")).toBeInTheDocument();
        expect(screen.getByText("Expense")).toBeInTheDocument();
        expect(screen.getByText("2023-01-01")).toBeInTheDocument();
        expect(screen.getByText("1234 ****")).toBeInTheDocument();
    });

    it("renders empty state correctly", () => {
        (useGetTransactionExpenseQuery as jest.Mock).mockReturnValue({
            data: { data: { content: [], totalPages: 0 } },
            error: null,
            isLoading: false,
        });

        render(<Expense />);
        expect(screen.getByText("No Expense found.")).toBeInTheDocument();
    });

    it("renders Pagination component when totalPages > 1", () => {
        const mockData = {
          data: {
            content: [],
            totalPages: 2,
          },
        };
    
        (useGetTransactionExpenseQuery as jest.Mock).mockReturnValue({
          data: mockData,
          error: null,
          isLoading: false,
        });
    
        render(<Expense />);
        expect(screen.getByRole("navigation")).toBeInTheDocument();
      });
})