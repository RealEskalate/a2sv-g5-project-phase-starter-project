import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import ActiveLoansOverviewTable from "../components/ActiveLoansOverviewTable/ActiveLoansOverviewTable";
import { useGetAllActiveLoansQuery } from "@/lib/redux/slices/activeLoanSlice";

// Mock the hook
jest.mock("../lib/redux/slices/activeLoanSlice", () => ({
  useGetAllActiveLoansQuery: jest.fn(),
}));

type MockedUseGetAllActiveLoansQuery = jest.MockedFunction<
  typeof useGetAllActiveLoansQuery
>;

describe("ActiveLoansOverviewTable Component", () => {
  const mockedUseGetAllActiveLoansQuery =
    useGetAllActiveLoansQuery as MockedUseGetAllActiveLoansQuery;

  beforeEach(() => {
    mockedUseGetAllActiveLoansQuery.mockReset();
  });

  it("should display data correctly when fetch is successful", async () => {
    const mockData = {
      data: [
        {
          serialNumber: "1",
          loanAmount: "$10,000",
          amountLeftToRepay: 5000,
          duration: 12,
          interestRate: 5,
          installment: 500,
          activeLoneStatus: "approved",
        },
        {
          serialNumber: "2",
          loanAmount: 20000,
          amountLeftToRepay: 15000,
          duration: 24,
          interestRate: 4.5,
          installment: 700,
          activeLoneStatus: "approved",
        },
      ],
      error: null,
      isLoading: false,
    };

    mockedUseGetAllActiveLoansQuery.mockImplementation(() => mockData as any);

    render(<ActiveLoansOverviewTable />);

    // Use regex or formatted strings to match the data rendered
    expect(await screen.findByText("$10,000")).toBeInTheDocument();
    expect(screen.getByText("$20,000")).toBeInTheDocument();
    expect(screen.getByText("$5,000")).toBeInTheDocument();
    expect(screen.getByText("$15,000")).toBeInTheDocument();
    expect(screen.getByText("Loan Money")).toBeInTheDocument();
    expect(screen.getByText("Left to repay")).toBeInTheDocument();
    expect(screen.getByText("Duration")).toBeInTheDocument();
    expect(screen.getByText("Interest rate")).toBeInTheDocument();
    expect(screen.getByText("Installment")).toBeInTheDocument();
    expect(screen.getByText("Repay")).toBeInTheDocument();
  });

  it("should handle loading state correctly", () => {
    mockedUseGetAllActiveLoansQuery.mockImplementation(
      () =>
        ({
          data: null,
          error: null,
          isLoading: true,
        } as any)
    );

    render(<ActiveLoansOverviewTable />);

    expect(screen.getByText("Loading...")).toBeInTheDocument(); // Adjust based on actual loading state display
  });

  it("should handle error state correctly", async () => {
    mockedUseGetAllActiveLoansQuery.mockImplementation(
      () =>
        ({
          data: null,
          error: "An error occurred",
          isLoading: false,
        } as any)
    );

    render(<ActiveLoansOverviewTable />);

    expect(screen.getByText("An error occurred")).toBeInTheDocument(); // Adjust based on actual error handling
  });
});
