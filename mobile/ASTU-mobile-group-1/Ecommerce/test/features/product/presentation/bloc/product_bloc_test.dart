import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/core/error/failure.dart';
import 'package:product_6/features/product/domain/entities/product_entity.dart';
import 'package:product_6/features/product/presentation/bloc/product_bloc.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockGetAllProductsUsecase mockGetAllProductsUsecase;
  late MockGetProductUsecase mockGetProductUsecase;
  late MockUpdateProductUsecase mockUpdateProductUsecase;
  late MockDeleteProductUsecase mockDeleteProductUsecase;
  late MockInsertProductUsecase mockInsertProductUsecase;
  late MockInputConverter mockInputConverter;

  late ProductBloc productBloc;
  setUp(() {
    mockGetAllProductsUsecase = MockGetAllProductsUsecase();
    mockGetProductUsecase = MockGetProductUsecase();
    mockUpdateProductUsecase = MockUpdateProductUsecase();
    mockInsertProductUsecase = MockInsertProductUsecase();
    mockDeleteProductUsecase = MockDeleteProductUsecase();
    mockInputConverter = MockInputConverter();

    productBloc = ProductBloc(
      getAllProductsUsecase: mockGetAllProductsUsecase,
      getProductUsecase: mockGetProductUsecase,
      updateProductUsecase: mockUpdateProductUsecase,
      deleteProductUsecase: mockDeleteProductUsecase,
      insertProductUsecase: mockInsertProductUsecase,
      inputConverter: mockInputConverter,
    );
  });

  const testProductEntity = ProductEntity(
    id: '1',
    name: 'shoe',
    description: 'best shoes',
    price: 12.0,
    imageUrl: 'http',
  );

  test('initial state should be empty', () {
    expect(productBloc.state, IntialState());
  });

  group(
    'LoadAllProductEvent',
    () {
      blocTest<ProductBloc, ProductState>(
        'emits [LoadingState, LoadedAllProductsState] when MyEvent is added.',
        build: () {
          when(mockGetAllProductsUsecase()).thenAnswer(
            (_) async => const Right(
              [testProductEntity],
            ),
          );
          return productBloc;
        },
        act: (bloc) => bloc.add(LoadAllProductEvent()),
        wait: const Duration(milliseconds: 500),
        expect: () => [
          LoadingState(),
          const LoadedAllProductsState(products: [testProductEntity])
        ],
      );

      blocTest<ProductBloc, ProductState>(
        'emits [LoadingState, ErrorState] when LoadAllProductEvent is added.',
        build: () {
          when(mockGetAllProductsUsecase()).thenAnswer(
            (_) async => const Left(ServerFailure('Server failure')),
          );
          return productBloc;
        },
        act: (bloc) => bloc.add(LoadAllProductEvent()),
        wait: const Duration(milliseconds: 500),
        expect: () => [
          LoadingState(),
          const ErrorState(message: 'Server failure'),
        ],
      );
    },
  );

  group('GetSingleProductEvent', () {
    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, LoadedSingleProductState] when GetSingleProductEvent is added.',
      build: () {
        when(mockGetProductUsecase('1')).thenAnswer(
          (_) async => const Right(
            testProductEntity,
          ),
        );
        return productBloc;
      },
      act: (bloc) => bloc.add(GetSingleProductEvent(id: '1')),
      expect: () => [
        LoadingState(),
        const LoadedSingleProductState(product: testProductEntity)
      ],
    );

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, ErrorState] when GetSingleProductEvent is added.',
      build: () {
        when(mockGetAllProductsUsecase()).thenAnswer(
          (_) async => const Left(ServerFailure('Server failure')),
        );
        return productBloc;
      },
      act: (bloc) => bloc.add(LoadAllProductEvent()),
      wait: const Duration(milliseconds: 500),
      expect: () => [
        LoadingState(),
        const ErrorState(message: 'Server failure'),
      ],
    );
  });
}
