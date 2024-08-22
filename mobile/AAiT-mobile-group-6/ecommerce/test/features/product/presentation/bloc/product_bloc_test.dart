import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failures.dart';
import 'package:ecommerce/features/product/data/model/product_model.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_event.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_state.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late MockGetProductUsecase mockGetProductUsecase;
  late MockGetAllProductUsecase mockGetAllProductUsecase;
  late MockDeleteProductUsecase mockDeleteProductUsecase;
  late MockInsertProductUsecase mockInsertProductUsecase;
  late MockUpdateProductUsecase mockUpdateProductUsecase;
  late ProductBloc productBloc;

  setUp(() {
    mockGetProductUsecase = MockGetProductUsecase();
    mockGetAllProductUsecase = MockGetAllProductUsecase();
    mockInsertProductUsecase = MockInsertProductUsecase();
    mockUpdateProductUsecase = MockUpdateProductUsecase();
    mockDeleteProductUsecase = MockDeleteProductUsecase();
    productBloc = ProductBloc(
        getProductUsecase: mockGetProductUsecase,
        getAllProductUsecase: mockGetAllProductUsecase,
        updateProductUsecase: mockUpdateProductUsecase,
        deleteProductUsecase: mockDeleteProductUsecase,
        insertProductUsecase: mockInsertProductUsecase);
  });

  final testId = '1';
  const testProduct = ProductModel(
    id: '1',
    name: 'name',
    description: 'description',
    price: 1.0,
    imageUrl: 'imageUrl',
  );

  test('initial state should be ProductStateEmpty', () {
    expect(productBloc.state, ProductStateEmpty());
  });

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, ProductStateLoaded] when successful',
    build: () {
      when(mockGetProductUsecase.execute(testId))
          .thenAnswer((_) async => const Right(testProduct));
      return productBloc;
    },
    act: (bloc) => bloc.add(GetProductEvent(productId: testId)),
    wait: const Duration(milliseconds: 500),
    expect: () =>
        [ProductStateLoading(), const ProductStateLoaded(product: testProduct)],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, ProductStateLoadedFailure] when unsuccessful',
    build: () {
      when(mockGetProductUsecase.execute(testId))
          .thenAnswer((_) async => Left(ServerFailure('Server Failure')));
      return productBloc;
    },
    act: (bloc) => bloc.add(GetProductEvent(productId: testId)),
    wait: const Duration(milliseconds: 500),
    expect: () => [
      ProductStateLoading(),
      const ProductLoadFailure(message: 'Server Failure')
    ],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, AllProductsLoaded] when successful',
    build: () {
      when(mockGetAllProductUsecase.execute(any))
          .thenAnswer((_) async => const Right([testProduct]));
      return productBloc;
    },
    act: (bloc) => bloc.add(GetAllProductEvent()),
    wait: const Duration(milliseconds: 10000),
    expect: () => [
      ProductStateLoading(),
      const AllProductsLoaded(products: [testProduct])
    ],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, AllProductsLoadedFailure] when unsuccessful',
    build: () {
      when(mockGetAllProductUsecase.execute(any))
          .thenAnswer((_) async => Left(ServerFailure('Server Failure')));
      return productBloc;
    },
    act: (bloc) => bloc.add(GetAllProductEvent()),
    wait: const Duration(milliseconds: 500),
    expect: () => [
      ProductStateLoading(),
      const AllProductsLoadedFailure(message: 'Server Failure')
    ],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, ProductInsertState] when successful',
    build: () {
      when(mockInsertProductUsecase.execute(testProduct))
          .thenAnswer((_) async => const Right(testProduct));
      return productBloc;
    },
    act: (bloc) => bloc.add(InsertProductEvent(product: testProduct)),
    wait: const Duration(milliseconds: 500),
    expect: () =>
        [ProductStateLoading(), const ProductInsertState(product: testProduct)],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, ProductInsertFailureState] when unsuccessful',
    build: () {
      when(mockInsertProductUsecase.execute(testProduct))
          .thenAnswer((_) async => Left(ServerFailure('Server Failure')));
      return productBloc;
    },
    act: (bloc) => bloc.add(InsertProductEvent(product: testProduct)),
    wait: const Duration(milliseconds: 500),
    expect: () => [
      ProductStateLoading(),
      const ProductInsertFailureState(message: 'Server Failure')
    ],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, ProductUpdateState] when successful',
    build: () {
      when(mockUpdateProductUsecase.execute(testProduct))
          .thenAnswer((_) async => const Right(testProduct));
      return productBloc;
    },
    act: (bloc) => bloc.add(UpdateProductEvent(product: testProduct)),
    wait: const Duration(milliseconds: 500),
    expect: () =>
        [ProductStateLoading(), const ProductUpdateState(product: testProduct)],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, ProductUpdateFailureState] when unsuccessful',
    build: () {
      when(mockUpdateProductUsecase.execute(testProduct))
          .thenAnswer((_) async => Left(ServerFailure('Server Failure')));
      return productBloc;
    },
    act: (bloc) => bloc.add(UpdateProductEvent(product: testProduct)),
    wait: const Duration(milliseconds: 500),
    expect: () => [
      ProductStateLoading(),
      const ProductUpdateFailureState(message: 'Server Failure')
    ],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, ProductDeleteState] when successful',
    build: () {
      when(mockDeleteProductUsecase.execute(testId))
          .thenAnswer((_) async => const Right(testProduct));
      return productBloc;
    },
    act: (bloc) => bloc.add(DeleteProductEvent(productId: testId)),
    wait: const Duration(milliseconds: 500),
    expect: () =>
        [ProductStateLoading(), const ProductDeleteState(product: testProduct)],
  );

  blocTest<ProductBloc, ProductState>(
    'should emit [ProductStateLoading, ProductUpdateFailureState] when unsuccessful',
    build: () {
      when(mockUpdateProductUsecase.execute(testProduct))
          .thenAnswer((_) async => Left(ServerFailure('Server Failure')));
      return productBloc;
    },
    act: (bloc) => bloc.add(UpdateProductEvent(product: testProduct)),
    wait: const Duration(milliseconds: 500),
    expect: () => [
      ProductStateLoading(),
      const ProductUpdateFailureState(message: 'Server Failure')
    ],
  );
}
