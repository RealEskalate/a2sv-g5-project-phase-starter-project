import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/usecase/usecase.dart';
import '../../domain/entities/product_entity.dart';
import '../../domain/usecases/add_product_usecase.dart';
import '../../domain/usecases/delete_product_usecase.dart';
import '../../domain/usecases/get_product_usecase.dart';
import '../../domain/usecases/get_products_usecase.dart';
import '../../domain/usecases/update_product_usecase.dart';

part 'product_event.dart';
part 'product_state.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState> {
  final GetProductsUsecase getProductsUsecase;
  final GetProductUsecase getProductUsecase;
  final AddProductUsecase addProductUsecase;
  final DeleteProductUsecase deleteProductUsecase;
  final UpdateProductUsecase updateProductUsecase;
  ProductBloc({
    required this.getProductsUsecase,
    required this.getProductUsecase,
    required this.addProductUsecase,
    required this.deleteProductUsecase,
    required this.updateProductUsecase,
  }) : super(ProductInitial()) {
    on<LoadAllProductEvent>((event, emit) async {
      emit(ProductLoading());

      final result = await getProductsUsecase(NoParams());
      result.fold(
        (failure) => emit(ProductErrorState(failure.message)),
        (products) => emit(LoadedAllProductState(products)),
      );
    });

    on<LoadSingleProductEvent>((event, emit) async {
      emit(ProductLoading());

      final result = await getProductUsecase(GetParams(id: event.id));

      result.fold(
        (failure) => emit(ProductErrorState(failure.message)),
        (product) => emit(LoadedSingleProductState(product)),
      );
    });

    on<DeleteProductEvent>(
      (event, emit) async {
        emit(ProductLoading());

        final result = await deleteProductUsecase(DeleteParams(id: event.id));

        result.fold(
            (failure) => emit(ProductErrorState(failure.message)),
            (value) => emit(
                const ProductDeletedState(message: 'Successfully Deleted Product')));
      },
    );

    on<UpdateProductEvent>((event, emit) async {
      emit(ProductLoading());
      final result =
          await updateProductUsecase(UpdateParams(product: event.product));
      result.fold(
        (failure) => emit(ProductErrorState(failure.message)),
        (product) => emit(ProductUpdatedState(event.product)),
      );
    });
    on<CreateProductEvent>(
      (event, emit) async {
        emit(ProductLoading());

        final result =
            await addProductUsecase(CreateParams(product: event.product));

        result.fold(
          (failure) => emit(ProductErrorState(failure.message)),
          (product) => emit(ProductCreatedState(event.product)),
        );
      },
    );
  }
}
