import 'package:application1/core/error/failure.dart';
import 'package:application1/core/usecase/usecase.dart';
import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/domain/usecases/add_product_usecase.dart';
import 'package:application1/features/product/domain/usecases/delete_product_usecase.dart';
import 'package:application1/features/product/domain/usecases/get_product_usecase.dart';
import 'package:application1/features/product/domain/usecases/update_product_usecase.dart';
import 'package:application1/features/product/presentation/bloc/product_bloc.dart';
import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';

void main() {
  late ProductBloc productBloc;
  late MockGetProductUsecase mockGetProductUsecase;
  late MockAddProductUsecase mockAddProductUsecase;
  late MockDeleteProductUsecase mockDeleteProductUsecase;
  late MockGetProductsUsecase mockGetProductsUsecase;
  late MockUpdateProductUsecase mockUpdateProductUsecase;
  setUp(() {
    mockUpdateProductUsecase = MockUpdateProductUsecase();
    mockGetProductsUsecase = MockGetProductsUsecase();
    mockGetProductUsecase = MockGetProductUsecase();
    mockDeleteProductUsecase = MockDeleteProductUsecase();
    mockUpdateProductUsecase = MockUpdateProductUsecase();
    mockAddProductUsecase = MockAddProductUsecase();
    productBloc = ProductBloc(
      getProductsUsecase: mockGetProductsUsecase,
      getProductUsecase: mockGetProductUsecase,
      addProductUsecase: mockAddProductUsecase,
      deleteProductUsecase: mockDeleteProductUsecase,
      updateProductUsecase: mockUpdateProductUsecase,
    );
  });

  const String id = '6672776eb905525c145fe0bb';
  const ProductEntity tProductEntity = ProductEntity(
    id: '6672776eb905525c145fe0bb',
    name: 'Anime website',
    description: 'Explore anime characters.',
    price: 123,
    imageUrl:
        'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg',
  );
   List<ProductEntity> tProductList = [
    const ProductEntity(
      id: '6672752cbd218790438efdb0',
      name: 'Anime website',
      description: 'Explore anime characters.',
      price: 123,
      imageUrl:
          'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg',
    ),
    const ProductEntity(
      id: '66728788b116d7a8d476558c',
      name: 'Better name',
      description: 'Even better description. The best description ever.',
      price: 112,
      imageUrl:
          'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718781833/images/jxt6xd4ivavvuodt9gkx.jpg',
    ),
  ];
  group('get product usecase', () {
    test('initial state must be empty', () async {
      expect(productBloc.state, ProductInitial());
    });

    blocTest<ProductBloc, ProductState>(
      'emits [ProductLoadingState] and then [LoadedSingleProductState] when LoadSingleProductEvent is added.',
      build: () {
        when(mockGetProductUsecase(const GetParams(id: id)))
            .thenAnswer((_) async => const Right(tProductEntity));
        return productBloc;
      },
      act: (bloc) => bloc.add(const LoadSingleProductEvent(id: id)),
      expect: () => [
        ProductLoading(),
        const LoadedSingleProductState(tProductEntity),
      ],
    );
    blocTest<ProductBloc, ProductState>(
      'should emit [ProductLoadingState,ProductErrorState] when get data is unsuccessful.',
      build: () {
        when(mockGetProductUsecase(const GetParams(id: id)))
            .thenAnswer((_) async => const Left(ServerFailure('server failure')));
        return productBloc;
      },
      act: (bloc) => bloc.add(const LoadSingleProductEvent(id: id)),
      expect: () => [
        ProductLoading(),
        const ProductErrorState('server failure'),
      ],
    );
  });

  group('get all products usecase', () {
    test('initial state must be empty', () async {
      expect(productBloc.state, ProductInitial());
    });

    blocTest<ProductBloc, ProductState>(
      'emits [ProductLoadingState,LoadedAllProductState] when LoadAllProductEvent is added.',
      build: () {
        when(mockGetProductsUsecase(NoParams()))
            .thenAnswer((_) async => Right(tProductList));
        return productBloc;
      },
      act: (bloc) => bloc.add(LoadAllProductEvent()),
      expect: () => [
        ProductLoading(),
        LoadedAllProductState(tProductList),
      ],
    );
    blocTest<ProductBloc, ProductState>(
      'should emit [ProductLoadingState,ProductErrorState] when get data is unsuccessful.',
      build: () {
        when(mockGetProductsUsecase(NoParams()))
            .thenAnswer((_) async => const Left(ServerFailure('server failure')));
        return productBloc;
      },
      act: (bloc) => bloc.add(LoadAllProductEvent()),
      expect: () => [
        ProductLoading(),
        const ProductErrorState('server failure'),
      ],
    );
  });
   group('add product usecase', () {
    test('initial state must be empty', () async {
      expect(productBloc.state, ProductInitial());
    });

    blocTest<ProductBloc, ProductState>(
      'emits [ProductLoadingState,ProductCreatedState] when CreateProductEvent is added.',
      build: () {
        when(mockAddProductUsecase(const CreateParams(product: tProductEntity)))
            .thenAnswer((_) async => const Right(tProductEntity));
        return productBloc;
      },
      act: (bloc) => bloc.add(const CreateProductEvent(product:  tProductEntity)),
      expect: () => [
        ProductLoading(),
        const ProductCreatedState(tProductEntity),
      ],
    );
    blocTest<ProductBloc, ProductState>(
      'should emit [ProductLoadingState,ProductErrorState] when add data is unsuccessful.',
      build: () {
        when(mockAddProductUsecase(const CreateParams(product: tProductEntity)))
            .thenAnswer((_) async => const Left(ServerFailure('server failure')));
        return productBloc;
      },
      act: (bloc) => bloc.add(const CreateProductEvent(product: tProductEntity)),
      expect: () => [
        ProductLoading(),
        const ProductErrorState('server failure'),
      ],
    );
  });

  group('update product usecase', () {
    test('initial state must be empty', () async {
      expect(productBloc.state, ProductInitial());
    });

    blocTest<ProductBloc, ProductState>(
      'emits [ProductLoadingState,ProductUpdatedState] when UpdateProductEvent is added.',
      build: () {
        when(mockUpdateProductUsecase(const UpdateParams(product: tProductEntity)))
            .thenAnswer((_) async => const Right(tProductEntity));
        return productBloc;
      },
      act: (bloc) => bloc.add(const UpdateProductEvent(product:  tProductEntity)),
      expect: () => [
        ProductLoading(),
        const ProductUpdatedState(tProductEntity),
      ],
    );
    blocTest<ProductBloc, ProductState>(
      'should emit [ProductLoadingState,ProductErrorState] when add data is unsuccessful.',
      build: () {
        when(mockUpdateProductUsecase(const UpdateParams(product: tProductEntity)))
            .thenAnswer((_) async => const Left(ServerFailure('server failure')));
        return productBloc;
      },
      act: (bloc) => bloc.add(const UpdateProductEvent(product:  tProductEntity)),
      expect: () => [
        ProductLoading(),
        const ProductErrorState('server failure'),
      ],
    );
  });

  group('delete product usecase', () {
    test('initial state must be empty', () async {
      expect(productBloc.state, ProductInitial());
    });

    blocTest<ProductBloc, ProductState>(
      'emits [ProductLoadingState,ProductDeletedState] when DeleteProductEvent is added.',
      build: () {
        when(mockDeleteProductUsecase(const DeleteParams(id: id)))
            .thenAnswer((_) async => const Right(true));
        return productBloc;
      },
      act: (bloc) => bloc.add(const DeleteProductEvent(id: id)),
      expect: () => [
        ProductLoading(),
        const ProductDeletedState( message: 'Successfully Deleted Product'),
      ],
    );
    blocTest<ProductBloc, ProductState>(
      'should emit [ProductLoadingState,ProductErrorState] when add data is unsuccessful.',
      build: () {
        when(mockDeleteProductUsecase(const DeleteParams(id:id)))
            .thenAnswer((_) async => const Left(ServerFailure('server failure')));
        return productBloc;
      },
      act: (bloc) => bloc.add(const DeleteProductEvent(id:id)),
      expect: () => [
        ProductLoading(),
        const ProductErrorState('server failure'),
      ],
    );
  });

}
