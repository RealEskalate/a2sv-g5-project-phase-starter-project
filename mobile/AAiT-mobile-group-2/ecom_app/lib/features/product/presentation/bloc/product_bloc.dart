import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/usecase/usecase.dart';
import '../../domain/entities/product.dart';
import '../../domain/usecases/create_product.dart';
import '../../domain/usecases/delete_product.dart';
import '../../domain/usecases/get_all_products.dart';
import '../../domain/usecases/get_current_product.dart';
import '../../domain/usecases/update_product.dart';

part 'product_event.dart';
part 'product_state.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState> {
  final GetAllProductsUsecase _getAllProductsUsecase;
  final GetCurrentProductUsecase _getCurrentProductUsecase;
  final CreateProductUsecase _createProductUsecase;
  final DeleteProductUsecase _deleteProductUsecase;
  final UpdateProductUsecase _updateProductUsecase;

  ProductBloc(
      this._getAllProductsUsecase,
      this._getCurrentProductUsecase,
      this._createProductUsecase,
      this._updateProductUsecase,
      this._deleteProductUsecase)
      : super(ProductInitialState()) {
    on<GetSingleProductEvent>((event, emit) async {
      emit(ProductLoading());
      final result = await _getCurrentProductUsecase(GetParams(id: event.id));
      result.fold(
          (failure) => emit(
                ProductErrorState(message: failure.message),
              ),
          (product) => emit(LoadSingleProductState(product: product)));
    });

    on<LoadAllProductEvent>((event, emit) async {
      print("LoadAllProductEvent has been dispatched");
      emit(ProductLoading());
      final result = await _getAllProductsUsecase(NoParams());
      result.fold(
          (failure) => emit(
                ProductErrorState(message: failure.message),
              ),
          (products) => emit(LoadAllProductState(products: products)));
    });

    on<CreateProductEvent>((event, emit) async {
      emit(ProductLoading());
      final result =
          await _createProductUsecase(CreateParams(product: event.product));
      result.fold(
          (failure) => emit(
                ProductCreatedErrorState(message: failure.message),
              ),
          (product) => emit(ProductCreatedState(product: product)));
    });

    on<UpdateProductEvent>((event, emit) async {
      emit(ProductLoading());
      final result =
          await _updateProductUsecase(UpdateParams(product: event.product));
      result.fold(
          (failure) => emit(
                ProductUpdatedErrorState(message: failure.message),
              ),
          (product) => emit(ProductUpdatedState(product: product)));
    });

    on<DeleteProductEvent>((event, emit) async {
      emit(ProductLoading());
      final result = await _deleteProductUsecase(DeleteParams(id: event.id));
      result.fold(
          (failure) => emit(
                ProductErrorState(message: failure.message),
              ),
          (product) => emit(ProductDeletedState()));
    });
  }
}
