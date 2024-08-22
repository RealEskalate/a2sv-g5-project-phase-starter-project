import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:rxdart/rxdart.dart';
import '../../domain/usecase/delete_product.dart';
import '../../domain/usecase/get_all_product.dart';
import '../../domain/usecase/get_product.dart';
import '../../domain/usecase/insert_product.dart';
import '../../domain/usecase/update_product.dart';
import 'product_event.dart';
import 'product_state.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState> {
  final GetProductUsecase getProductUsecase;
  final GetAllProductUsecase getAllProductUsecase;
  final UpdateProductUsecase updateProductUsecase;
  final DeleteProductUsecase deleteProductUsecase;
  final InsertProductUsecase insertProductUsecase;

  ProductBloc({
    required this.getProductUsecase,
    required this.getAllProductUsecase,
    required this.updateProductUsecase,
    required this.deleteProductUsecase,
    required this.insertProductUsecase,
  }) : super(ProductStateEmpty()) {
    on<GetProductEvent>((event, emit) async {
      emit(ProductStateLoading());
      final result = await getProductUsecase.execute(event.productId);
      result.fold(
          (failure) => emit(ProductLoadFailure(message: failure.message)),
          (product) => emit(ProductStateLoaded(product: product)));
    }, transformer: debounce(const Duration(milliseconds: 500)));

    on<GetAllProductEvent>((event, emit) async {
      emit(ProductStateLoading());
      final result = await getAllProductUsecase.execute(NoParams());
      result.fold(
          (failure) => emit(AllProductsLoadedFailure(message: failure.message)),
          (products) => emit(AllProductsLoaded(products: products)));
    }, transformer: debounce(const Duration(milliseconds: 500)));
    on<InsertProductEvent>((event, emit) async {
      emit(ProductStateLoading());
      final result = await insertProductUsecase.execute(event.product);
      result.fold(
          (failure) =>
              emit(ProductInsertFailureState(message: failure.message)),
          (product) => emit(ProductInsertState(product: product)));
    }, transformer: debounce(const Duration(milliseconds: 500)));
    on<UpdateProductEvent>((event, emit) async {
      emit(ProductStateLoading());
      final result = await updateProductUsecase.execute(event.product);
      result.fold(
          (failure) =>
              emit(ProductUpdateFailureState(message: failure.message)),
          (product) => emit(ProductUpdateState(product: product)));
    }, transformer: debounce(const Duration(milliseconds: 500)));
    on<DeleteProductEvent>((event, emit) async {
      emit(ProductStateLoading());
      final result = await deleteProductUsecase.execute(event.productId);
      result.fold(
          (failure) =>
              emit(ProductDeleteFailureState(message: failure.message)),
          (product) => emit(ProductDeleteState(product: product)));
    }, transformer: debounce(const Duration(milliseconds: 500)));
  }
}

EventTransformer<T> debounce<T>(Duration duration) {
  return (events, mapper) {
    return events.debounceTime(duration).flatMap(mapper);
  };
}
