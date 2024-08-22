import 'package:bloc/bloc.dart';
import '../../../../core/usecase/usecase.dart';
import '../../domain/use_case/delete_product_usecase.dart';
import '../../domain/use_case/get_product_by_id_usecase.dart';
import '../../domain/use_case/get_products_usecase.dart';
import '../../domain/use_case/insert_product_usecase.dart';
import '../../domain/use_case/update_product_usecase.dart';
import 'product_event.dart';
import 'product_state.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState> {
  final GetProductsUsecase getProductsUsecase;
  final GetProductByIdUsecase getProductByIdUsecase;
  final InsertProductUsecase insertProductUsecase;
  final UpdateProductUsecase updateProductUsecase;
  final DeleteProductUsecase deleteProductUsecase;

  ProductBloc({
    required this.getProductsUsecase,
    required this.getProductByIdUsecase,
    required this.insertProductUsecase,
    required this.updateProductUsecase,
    required this.deleteProductUsecase,
  }) : super(ProductInitial()) {
    
    on<LoadProduct>((event, emit) async {
      emit(ProductLoading());
      final result = await getProductsUsecase.call(NoParams());
      result.fold(
        (failure) => emit(ProductError()),
        (products) => emit(ProductLoaded(products: products)),
      );
    });

    on<GetProductByIdEvent>((event, emit) async {
      emit(ProductLoading());
      final result = await getProductByIdUsecase(GetParams(id: event.id));
      result.fold(
        (failure) => emit(ProductError()),
        (product) => emit(ProductLoadedSingle(product: product)),
      );
    });

    on<CreateProductEvent>((event, emit) async {
      emit(ProductLoading());
      final result = await insertProductUsecase.call(CreateParams(product: event.product));
      result.fold(
        (failure) => emit(const ProductCreatedErrorState(message: 'Error creating product')),
        (product) => emit(ProductCreatedState(product: product)),
      );
    });


    on<UpdateProductEvent>((event, emit) async {
      emit(ProductLoading());
      final result = await updateProductUsecase.call(UpdateParams(product: event.product));
      result.fold(
        (failure) => emit(const ProductUpdatedErrorState(message: 'Error updating product')),
        (product) => emit(ProductUpdatedState(product: product)),
      );
    });

    on<DeleteProductEvent>((event, emit) async {
      emit(ProductLoading());
      final result = await deleteProductUsecase.call(DeleteParams(id: event.id));
      result.fold(
        (failure) => emit(ProductError()),
        (product) => emit(ProductDeleteState()),
      );
    });
    
  }
}
