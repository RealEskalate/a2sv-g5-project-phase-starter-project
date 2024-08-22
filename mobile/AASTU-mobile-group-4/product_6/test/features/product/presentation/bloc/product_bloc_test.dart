// product_bloc_test.dart

import 'dart:ffi';

import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/features/product/domain/entities/product.dart';
import 'package:product_6/features/product/presentation/bloc/product_bloc.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:bloc_test/bloc_test.dart';
import 'package:product_6/features/product/presentation/bloc/product_event.dart';
import 'package:product_6/features/product/presentation/bloc/product_state.dart';

import '../../../../mock.mocks.dart';


// Mocking Use Cases
// class MockCreateProductUseCase extends Mock implements CreateProductUseCase {}
// class MockDeleteProductUseCase extends Mock implements DeleteProductUseCase {}
// class MockUpdateProductUseCase extends Mock implements UpdateProductUseCase {}
// class MockViewAllProductsUseCase extends Mock implements ViewAllProductsUseCase {}
// class MockViewProductUseCase extends Mock implements ViewProductUseCase {}

void main() {
  late ProductBloc productBloc;
  late MockCreateProductUseCase mockCreateProductUseCase;
  late MockDeleteProductUseCase mockDeleteProductUseCase;
  late MockUpdateProductUseCase mockUpdateProductUseCase;
  late MockViewAllProductsUseCase mockViewAllProductsUseCase;
  late MockViewProductUseCase mockViewProductUseCase;

  setUp(() {
    mockCreateProductUseCase = MockCreateProductUseCase();
    mockDeleteProductUseCase = MockDeleteProductUseCase();
    mockUpdateProductUseCase = MockUpdateProductUseCase();
    mockViewAllProductsUseCase = MockViewAllProductsUseCase();
    mockViewProductUseCase = MockViewProductUseCase();

    productBloc = ProductBloc(
      createProductUseCase: mockCreateProductUseCase,
      deleteProductUseCase: mockDeleteProductUseCase,
      updateProductUseCase: mockUpdateProductUseCase,
      viewAllProductsUseCase: mockViewAllProductsUseCase,
      viewProductUseCase: mockViewProductUseCase,
    );
  });

  tearDown(() {
    productBloc.close();
  });

  group('LoadAllProductEvent', () {
    final List<Product> products = [
      Product(id: '1', name: 'Product 1', description: 'Description 1', imageUrl: '', price: 0.0),
      Product(id: '2', name: 'Product 2', description: 'Description 2', imageUrl: '', price: 0.0),
    ];

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, LoadedAllProductState] when LoadAllProductEvent is added and succeeds',
      build: () {
        when(mockViewAllProductsUseCase())
            .thenAnswer((_) async => Right(products));
        return productBloc;
      },
      act: (bloc) => bloc.add(LoadAllProductEvent()),
      expect: () => [
        LoadingState(),
        LoadedAllProductState(products),
      ],
      verify: (_) {
        verify(mockViewAllProductsUseCase()).called(1);
      },
    );

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, ErrorState] when LoadAllProductEvent is added and fails',
      build: () {
        when(mockViewAllProductsUseCase())
            .thenAnswer((_) async => Left(ServerFailure('An error has occured')));
        return productBloc;
      },
      act: (bloc) => bloc.add(LoadAllProductEvent()),
      expect: () => [
        LoadingState(),
        ErrorState('An error occurred'),
      ],
    );
  });

  // group('GetSingleProductEvent', () {
    final Product product = Product(id: '1', name: 'Product 1', description: 'Description 1', imageUrl: '', price: 0.0);

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, LoadedSingleProductState] when GetSingleProductEvent is added and succeeds',
      build: () {
        when(mockViewProductUseCase(any))
            .thenAnswer((_) async => Right(product));
        return productBloc;
      },
      act: (bloc) => bloc.add(GetSingleProductEvent( '1')),
      expect: () => [
        LoadingState(),
        LoadedSingleProductState(product),
      ],
    );

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, ErrorState] when GetSingleProductEvent is added and fails',
      build: () {
        when(mockViewProductUseCase(any))
            .thenAnswer((_) async => Left(ServerFailure()));
        return productBloc;
      },
      act: (bloc) => bloc.add(GetSingleProductEvent('1')),
      expect: () => [
        LoadingState(),
        ErrorState('An error occurred'),
      ],
    );
  });

  // group('CreateProductEvent', () {
    final Product product = Product(id: '1', name: 'New Product', description: 'New Description');

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, LoadedSingleProductState] when CreateProductEvent is added and succeeds',
      build: () {
        when(mockCreateProductUseCase(any))
            .thenAnswer((_) async => Right(product));
        return productBloc;
      },
      act: (bloc) => bloc.add(CreateProductEvent(product: product, id: '')),
      expect: () => [
        LoadingState(),
        LoadedSingleProductState(product),
      ],
    );

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, ErrorState] when CreateProductEvent is added and fails',
      build: () {
        when(mockCreateProductUseCase(any))
            .thenAnswer((_) async => Left(ServerFailure()));
        return productBloc;
      },
      act: (bloc) => bloc.add(CreateProductEvent(product: product)),
      expect: () => [
        LoadingState(),
        ErrorState('An error occurred'),
      ],
    );
  });

  // group('UpdateProductEvent', () {
    final Product product = Product(id: '1', name: 'Updated Product', description: 'Updated Description');

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, LoadedSingleProductState] when UpdateProductEvent is added and succeeds',
      build: () {
        when(mockUpdateProductUseCase(any))
            .thenAnswer((_) async => Right(product));
        return productBloc;
      },
      act: (bloc) => bloc.add(UpdateProductEvent(product: product)),
      expect: () => [
        LoadingState(),
        LoadedSingleProductState(product),
      ],
    );

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, ErrorState] when UpdateProductEvent is added and fails',
      build: () {
        when(mockUpdateProductUseCase(any))
            .thenAnswer((_) async => Left(ServerFailure()));
        return productBloc;
      },
      act: (bloc) => bloc.add(UpdateProductEvent(product: product)),
      expect: () => [
        LoadingState(),
        ErrorState('An error occurred'),
      ],
    );
  });

  // group('DeleteProductEvent', () {
    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, InitialState] when DeleteProductEvent is added and succeeds',
      build: () {
        when(mockDeleteProductUseCase(any))
            .thenAnswer((_) async => Right(null));
        return productBloc;
      },
      act: (bloc) => bloc.add(DeleteProductEvent(productId: '1')),
      expect: () => [
        LoadingState(),
        InitialState(),
      ],
    );

    blocTest<ProductBloc, ProductState>(
      'emits [LoadingState, ErrorState] when DeleteProductEvent is added and fails',
      build: () {
        when(mockDeleteProductUseCase(any))
            .thenAnswer((_) async => Left(ServerFailure()));
        return productBloc;
      },
      act: (bloc) => bloc.add(DeleteProductEvent(productId: '1')),
      expect: () => [
        LoadingState(),
        ErrorState('An error occurred'),
      ],
    );
  });
