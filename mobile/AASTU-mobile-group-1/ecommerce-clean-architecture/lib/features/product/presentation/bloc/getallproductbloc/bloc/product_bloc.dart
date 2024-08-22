// import 'dart:core';
import 'dart:math';

import 'package:bloc/bloc.dart';

import 'package:equatable/equatable.dart';

import '../../../../../../core/error/failure.dart';
import '../../../../data/model/product_model.dart';
import '../../../../domain/usecases/addproduct.dart';
import '../../../../domain/usecases/deleteproduct.dart';
import '../../../../domain/usecases/getallproduct.dart';
import '../../../../domain/usecases/getproduct.dart';
import '../../../../domain/usecases/updateproduct.dart';

part 'product_event.dart';
part 'product_state.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState> {
  final GetAllProductUsecase getAllProductUsecase;
  final AddProductUsecase addProductUsecase;
  final UpdateProductUsecase updateProductUsecase;
  final DeleteProductUsecase deleteProductUsecase;
  final GetProductUsecase getProductUsecase;

  ProductBloc({
    required this.getAllProductUsecase,
    required this.addProductUsecase,
    required this.updateProductUsecase,
    required this.deleteProductUsecase,
    required this.getProductUsecase,
  }) : super(homeloading()) {

    on<GetAllProductEvent>(_onGetAllProducts);
    on<AddProductEvent>(_onAddProduct);
    on<UpdateProductEvent>(_onUpdateProduct);
    on<DeleteProductEvent>(_onDeleteProduct);
    on<GetProductEvent>(_onGetProduct);
  }
Future<void> _onGetAllProducts(GetAllProductEvent event, Emitter<ProductState> emit) async {
  emit(homeloading());
  var products = await getAllProductUsecase.getall();
  products.fold(
  (failure) => emit(homefailure(failure.message)), (products) => emit(homeloaded(products.map((e) => ProductModel.fromEntity(e)).toList()))
  );
}

Future<void>_onAddProduct(AddProductEvent event, Emitter<ProductState> emit)async{
  emit(adding());

   var productResult = await addProductUsecase.add(event.product);
  // print(productResult);
  productResult.fold(
      (failure) => emit(addfailure(failure.message)),
      (_) => emit(added()),
    );

}
Future<void>_onDeleteProduct(DeleteProductEvent event, Emitter<ProductState> emit)async{
  emit(deleting());
  var productResult = await deleteProductUsecase.delete(event.id);
  productResult.fold(
      (failure) => emit(deletefailure(failure.message)),
      (_) => emit(deleted()),
    );
}
Future <void> _onUpdateProduct(UpdateProductEvent event, Emitter<ProductState> emit) async {
  emit(updating());
  var productResult = await updateProductUsecase.update(event.product);
  // print("rediet");
  // print(state);
  // print(productResult);
  productResult.fold(
      (failure) => emit(updatefailure(failure.message)),
      (_) => emit(updated()),
    );
}
Future<void> _onGetProduct(GetProductEvent event, Emitter<ProductState> emit) async {
  emit(getloading());
  var product = await getProductUsecase.getprod(event.id);
  print(product);
  print(state);
  product.fold(
    (failure) => emit(getfailure(failure.message)),
    (product) => emit(getloaded(ProductModel.fromEntity(product)))
  );
}


}




























