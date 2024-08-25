import 'package:bloc/bloc.dart';

import '../../domain/usecase/add_product.dart';
import '../../domain/usecase/delete_product.dart';
import '../../domain/usecase/get_all_product.dart';
import '../../domain/usecase/get_product.dart';
import '../../domain/usecase/update_product.dart';
import 'events.dart';
import 'states.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState> {
  final GetAllProductUseCase getAllProductUseCase;
  final GetProductUseCase getProductUseCase;
  final UpdateProductUseCase updateProductUseCase;
  final DeleteProductUseCase deleteProductUseCase;
  final AddProductUseCase addProductUseCase;

  ProductBloc({
     required this.getAllProductUseCase,
      required this.getProductUseCase,
    required this.updateProductUseCase,
    required this.deleteProductUseCase,
    required this.addProductUseCase,
  }) : super(InitialState()){
    on<GetAllProductEvent>((event, emit) async {
      emit( LoadingState());
      try {
        print("event is GetAllProductEvent");
        // print('Getting all products');
        final products = await getAllProductUseCase.execute();
        print("$products is the products");
        products.fold((l) => emit(ErrorState(l.message)), (r) => emit(LoadedState(r)));
      } catch (e) {
        emit(ErrorState('Failed to get all products'));
      }
    });
    on<GetProductEvent>((event, emit) async {
      emit(LoadingState());
      try {
        final product = await getProductUseCase.call(event.productId);
        product.fold((failure) => emit(ErrorState('Failed to get Product')), (product) => emit(GetProductState(product)));
      } catch (e) {
        emit(ErrorState('Failed to get the product'));
      }
    });
    on<UpdateProductEvent>((event, emit) async {
      emit(LoadingState());
      try {
        await updateProductUseCase.call(id: event.productId, name: event.newName, price: event.newPrice, description: event.newDescription);
        emit(SuccessState('Product updated successfully'));
      } catch (e) {
        emit(ErrorState('Failed to update the product'));
      }
    });
    on<DeleteProductEvent>((event, emit) async {
      emit(LoadingState());
      try {
        await deleteProductUseCase.call(event.productId);
        emit(SuccessState('Product deleted successfully'));
      } catch (e) {
        emit(ErrorState('Failed to delete the product'));
      }
    });
    on<AddProductEvent>((event, emit) async {
      emit(InitialState());
      try {
        await addProductUseCase.call(event.product);
        emit(AddProductState(event.product));
      } catch (e) {
        emit(ErrorState('Failed to create the product'));
      }
    });
    @override
    ProductBloc copyWith({
      GetAllProductUseCase? getAllProductUseCase,
      GetProductUseCase? getProductUseCase,
      UpdateProductUseCase? updateProductUseCase,
      DeleteProductUseCase? deleteProductUseCase,
      AddProductUseCase? addProductUseCase,
    }) {
      return ProductBloc(
        getAllProductUseCase: getAllProductUseCase ?? this.getAllProductUseCase,
        getProductUseCase: getProductUseCase ?? this.getProductUseCase,
        updateProductUseCase: updateProductUseCase ?? this.updateProductUseCase,
        deleteProductUseCase: deleteProductUseCase ?? this.deleteProductUseCase,
        addProductUseCase: addProductUseCase ?? this.addProductUseCase,
      );
    }
  }

}