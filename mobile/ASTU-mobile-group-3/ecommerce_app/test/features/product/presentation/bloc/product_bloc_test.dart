import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/core/constants/constants.dart';
import 'package:ecommerce_app/core/errors/failures/failure.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_events.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_states.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main() {
  late MockGetAllProductUseCase mockGetAllProductUseCase;
  late MockUpdateProductUsecase mockUpdateProductUsecase;
  late MockDeleteProductUseCase mockDeleteProductUseCase;
  late MockInsertProductUseCase mockInsertProductUseCase;
  late MockGetProductUseCase mockGetProductUseCase;
  late ProductBloc productBloc;

  setUp(() {
    mockGetProductUseCase = MockGetProductUseCase();
    mockGetAllProductUseCase = MockGetAllProductUseCase();
    mockUpdateProductUsecase = MockUpdateProductUsecase();
    mockInsertProductUseCase = MockInsertProductUseCase();
    mockDeleteProductUseCase = MockDeleteProductUseCase();
    productBloc = ProductBloc(
      getAllProductUseCase: mockGetAllProductUseCase,
      deleteProductUseCase: mockDeleteProductUseCase,
      getProductUseCase: mockGetProductUseCase,
      insertProductUseCase: mockInsertProductUseCase,
      updateProductUsecase: mockUpdateProductUsecase,
    );
  });

  group('Bloc initial state check', () {
    test('Initial State should be InitialStateInstance', () {
      /// assert
      expect(productBloc.state, InitialState());
    });
  });

  group('getProducts test', () {
    blocTest<ProductBloc, ProductStates>(
      'Should return actual value when succes',
      build: () {
        when(mockGetProductUseCase.execute(any))
            .thenAnswer((_) async => Right(TestingDatas.testDataEntity));
        return productBloc;
      },
      act: (bloc) => bloc.add(GetSingleProductEvents(id: TestingDatas.id)),
      expect: () => [
        LoadingState(),
        LoadedSingleProductState(productEntity: TestingDatas.testDataEntity)
      ],
    );

    blocTest<ProductBloc, ProductStates>(
      'Should return server error state  when failed',
      build: () {
        when(mockGetProductUseCase.execute(any)).thenAnswer((_) async =>
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(GetSingleProductEvents(id: TestingDatas.id)),
      expect: () => [
        LoadingState(),
        ErrorState(message: AppData.getMessage(AppData.serverError)),
      ],
    );

    blocTest<ProductBloc, ProductStates>(
      'Should return cache error when no network',
      build: () {
        when(mockGetProductUseCase.execute(any)).thenAnswer((_) async =>
            Left(CacheFailure(AppData.getMessage(AppData.cacheError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(GetSingleProductEvents(id: TestingDatas.id)),
      expect: () => [
        LoadingState(),
        ErrorState(message: AppData.getMessage(AppData.cacheError)),
      ],
    );
  });

  group('getAllProduct test', () {
    blocTest<ProductBloc, ProductStates>(
      'Should return actual value when succes',
      build: () {
        when(mockGetAllProductUseCase.execute())
            .thenAnswer((_) async => Right(TestingDatas.productEntityList));
        return productBloc;
      },
      act: (bloc) => bloc.add(LoadAllProductEvents()),
      expect: () => [
        LoadingState(),
        LoadedAllProductState(data: TestingDatas.productEntityList)
      ],
    );

    blocTest<ProductBloc, ProductStates>(
      'Should return Server error  when failed',
      build: () {
        when(mockGetAllProductUseCase.execute()).thenAnswer((_) async =>
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(LoadAllProductEvents()),
      expect: () => [
        LoadingState(),
        ErrorState(message: AppData.getMessage(AppData.serverError))
      ],
    );

    blocTest<ProductBloc, ProductStates>(
      'Should return cache error when there is caching problem',
      build: () {
        when(mockGetAllProductUseCase.execute()).thenAnswer((_) async =>
            Left(CacheFailure(AppData.getMessage(AppData.cacheError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(LoadAllProductEvents()),
      expect: () => [
        LoadingState(),
        ErrorState(message: AppData.getMessage(AppData.cacheError))
      ],
    );
  });

  group('insert Product test', () {
    blocTest<ProductBloc, ProductStates>(
      'Should return int updated success code  when success',
      build: () {
        when(mockInsertProductUseCase.execute(any))
            .thenAnswer((_) async => const Right(AppData.successInsert));
        return productBloc;
      },
      act: (bloc) => bloc.add(InsertProductEvent(
          name: TestingDatas.testDataEntity.name,
          imageUrl: File(TestingDatas.testDataEntity.imageUrl),
          price: TestingDatas.testDataEntity.price.toString(),
          description: TestingDatas.testDataEntity.description)),
      expect: () => [
        LoadingState(),
        SuccessfullState(message: AppData.message[AppData.successInsert]!)
      ],
    );
    blocTest<ProductBloc, ProductStates>(
      'Should return int 1 when success',
      build: () {
        when(mockInsertProductUseCase.execute(any)).thenAnswer((_) async =>
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(InsertProductEvent(
          name: TestingDatas.testDataEntity.name,
          imageUrl: File(TestingDatas.testDataEntity.imageUrl),
          price: TestingDatas.testDataEntity.price.toString(),
          description: TestingDatas.testDataEntity.description)),
      expect: () => [
        LoadingState(),
        ErrorState(
          message: AppData.getMessage(AppData.serverError),
        ),
      ],
    );
    blocTest<ProductBloc, ProductStates>(
      'Should return  connection error when no network',
      build: () {
        when(mockInsertProductUseCase.execute(any)).thenAnswer((_) async =>
            Left(ConnectionFailure(
                AppData.getMessage(AppData.connectionError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(InsertProductEvent(
          name: TestingDatas.testDataEntity.name,
          imageUrl: File(TestingDatas.testDataEntity.imageUrl),
          price: TestingDatas.testDataEntity.price.toString(),
          description: TestingDatas.testDataEntity.description)),
      expect: () => [
        LoadingState(),
        ErrorState(
          message: AppData.getMessage(AppData.connectionError),
        ),
      ],
    );
  });

  group('update Product test', () {
    blocTest<ProductBloc, ProductStates>(
      'Should update and return  update success code  when success',
      build: () {
        when(mockUpdateProductUsecase.execute(any))
            .thenAnswer((_) async => const Right(AppData.successUpdate));
        return productBloc;
      },
      act: (bloc) => bloc.add(UpdateProductEvent(
          id: TestingDatas.testDataEntity.id,
          name: TestingDatas.testDataEntity.name,
          price: TestingDatas.testDataEntity.price.toString(),
          description: TestingDatas.testDataEntity.description)),
      expect: () => [
        LoadingState(),
        SuccessfullState(message: AppData.message[AppData.successUpdate]!)
      ],
    );
    blocTest<ProductBloc, ProductStates>(
      'Should return server failure message when there is no connection',
      build: () {
        when(mockUpdateProductUsecase.execute(any)).thenAnswer((_) async =>
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(UpdateProductEvent(
          id: TestingDatas.testDataEntity.id,
          name: TestingDatas.testDataEntity.name,
          price: TestingDatas.testDataEntity.price.toString(),
          description: TestingDatas.testDataEntity.description)),
      expect: () => [
        LoadingState(),
        ErrorState(
          message: AppData.getMessage(AppData.serverError),
        ),
      ],
    );
    blocTest<ProductBloc, ProductStates>(
      'Should return  connection error when no network',
      build: () {
        when(mockUpdateProductUsecase.execute(any)).thenAnswer((_) async =>
            Left(ConnectionFailure(
                AppData.getMessage(AppData.connectionError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(UpdateProductEvent(
          id: TestingDatas.testDataEntity.id,
          name: TestingDatas.testDataEntity.name,
          price: TestingDatas.testDataEntity.price.toString(),
          description: TestingDatas.testDataEntity.description)),
      expect: () => [
        LoadingState(),
        ErrorState(
          message: AppData.getMessage(AppData.connectionError),
        ),
      ],
    );
  });

  group('delte Product test', () {
    blocTest<ProductBloc, ProductStates>(
      'Should delete and return  update success code  when success',
      build: () {
        when(mockDeleteProductUseCase.execute(any))
            .thenAnswer((_) async => const Right(AppData.successDelete));
        return productBloc;
      },
      act: (bloc) => bloc.add(DeleteProductEvent(
        id: TestingDatas.testDataEntity.id,
      )),
      expect: () => [
        LoadingState(),
        SuccessfullState(message: AppData.message[AppData.successDelete]!)
      ],
    );
    blocTest<ProductBloc, ProductStates>(
      'Should return server failure message when there is no connection',
      build: () {
        when(mockDeleteProductUseCase.execute(any)).thenAnswer((_) async =>
            Left(ServerFailure(AppData.getMessage(AppData.serverError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(DeleteProductEvent(
        id: TestingDatas.testDataEntity.id,
      )),
      expect: () => [
        LoadingState(),
        ErrorState(
          message: AppData.getMessage(AppData.serverError),
        ),
      ],
    );
    blocTest<ProductBloc, ProductStates>(
      'Should return  connection error when no network',
      build: () {
        when(mockDeleteProductUseCase.execute(any)).thenAnswer((_) async =>
            Left(ConnectionFailure(
                AppData.getMessage(AppData.connectionError))));
        return productBloc;
      },
      act: (bloc) => bloc.add(DeleteProductEvent(
        id: TestingDatas.testDataEntity.id,
      )),
      expect: () => [
        LoadingState(),
        ErrorState(
          message: AppData.getMessage(AppData.connectionError),
        ),
      ],
    );
  });
}
