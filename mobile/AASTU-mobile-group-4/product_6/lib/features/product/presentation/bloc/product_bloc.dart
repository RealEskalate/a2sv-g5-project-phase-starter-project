import 'package:bloc/bloc.dart';
import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../../domain/entities/product.dart';
import '../../domain/usecases/create_product.dart';
import '../../domain/usecases/delete_product.dart';
import '../../domain/usecases/update_product.dart';
import '../../domain/usecases/view_all_products.dart';
import '../../domain/usecases/view_product.dart';
import 'product_event.dart';
import 'product_state.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState> {
  final CreateProductUseCase createProductUseCase;
  final DeleteProductUseCase deleteProductUseCase;
  final UpdateProductUseCase updateProductUseCase;
  final ViewAllProductsUseCase viewAllProductsUseCase;
  final ViewProductUseCase viewProductUseCase;

  ProductBloc({
    required this.createProductUseCase,
    required this.deleteProductUseCase,
    required this.updateProductUseCase,
    required this.viewAllProductsUseCase,
    required this.viewProductUseCase,
  }) : super(InitialState()) {
    on<LoadAllProductEvent>(_onLoadAllProducts);
    on<GetSingleProductEvent>(_onGetSingleProduct);
    on<UpdateProductEvent>(_onUpdateProduct);
    on<DeleteProductEvent>(_onDeleteProduct);
    on<CreateProductEvent>(_onCreateProduct);
  }

  Future<void> _onLoadAllProducts(
    LoadAllProductEvent event,
    Emitter<ProductState> emit,
  ) async {
    emit(LoadingState());

    final failureOrProducts = await viewAllProductsUseCase();
    print(failureOrProducts);
    emit(failureOrProducts.fold(
      (failure) => ErrorState(_mapFailureToMessage(failure)),
      (products) => LoadedAllProductState(products),
    ));
  }

  Future<void> _onGetSingleProduct(
    GetSingleProductEvent event,
    Emitter<ProductState> emit,
  ) async {
    emit(LoadingState());

    final Either<Failure, Product> failureOrProduct =
        await viewProductUseCase(event.productId);
    print('from single bloc');
    print(failureOrProduct);

    emit(failureOrProduct.fold(
      (failure) => ErrorState(_mapFailureToMessage(failure)),
      (product) => LoadedSingleProductState(product: product),
    ));
  }

  Future<void> _onUpdateProduct(
    UpdateProductEvent event,
    Emitter<ProductState> emit,
  ) async {
    emit(LoadingState());

    final Either<Failure, Product> failureOrProduct =
        await updateProductUseCase(event.product);
    print(failureOrProduct);
    emit(failureOrProduct.fold(
      (failure) => ErrorState(_mapFailureToMessage(failure)),
      (product) => LoadedSingleProductState(product: product),
    ));
  }

  Future<void> _onDeleteProduct(
    DeleteProductEvent event,
    Emitter<ProductState> emit,
  ) async {
    emit(LoadingState());

    final Either<Failure, void> failureOrSuccess =
        await deleteProductUseCase(event.productId);
    print("suuuuuuuuu ${failureOrSuccess}");
    emit(failureOrSuccess.fold(
      (failure) => ErrorState(_mapFailureToMessage(failure)),
      (_) => Success(
          message:
              "product deleted sucessfully"), // Assuming that after deletion, you return to the initial state
    ));
  }

  Future<void> _onCreateProduct(
    CreateProductEvent event,
    Emitter<ProductState> emit,
  ) async {
    emit(LoadingState());

    final failureOrProduct = await createProductUseCase(event.product);
    print(failureOrProduct);
    emit(failureOrProduct.fold(
      (failure) => ErrorState(_mapFailureToMessage(failure)),
      (product) => LoadedSingleProductState(product: product),
    ));
  }

  String _mapFailureToMessage(Failure failure) {
    // Customize the error message handling based on your Failure classes
    return 'An error occurred'; // A placeholder, replace it with your logic
  }
}
