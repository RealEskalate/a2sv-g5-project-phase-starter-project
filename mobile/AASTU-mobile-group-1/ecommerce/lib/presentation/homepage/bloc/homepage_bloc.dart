import 'dart:async';

import 'package:meta/meta.dart';
import 'package:ecommerce/core/import/import_file.dart';

part 'homepage_event.dart';
part 'homepage_state.dart';

class HomePageBloc extends Bloc<HomepageEvent, HomepageState> {
  GetAllProductUsecase getAllProductUsecase;
  DeleteProductUsecase deleteProductUsecase;
  AddProductUsecase addProductUsecase;
  UpdateProductUsecase updateProductUsecase;
  HomePageBloc({
    required this.updateProductUsecase,
    required this.addProductUsecase,
    required this.deleteProductUsecase,
    required this.getAllProductUsecase,
  }) : super(HomepageState()) {
     on<FetchProducts>(_fetchAllProduct);
    on<DeleteButtonPress>(_deleteProduct);
    on<AddButtonPress>(_addProduct);
    on<UpdateButtonPress>(_updateProduct);
  }

  Future<void> _fetchAllProduct(FetchProducts event, Emitter<HomepageState> emit) async{
 
    emit(state.copyWith(status: HomepageStatus.loading));
    try {
      final stream = getAllProductUsecase.execute();
      await for (final result in stream) {
        emit(state.copyWith(status: HomepageStatus.loaded, products: result));
      }
    } catch (e) {
      emit(state.copyWith(status: HomepageStatus.error));
    }
  }
  FutureOr<void> _deleteProduct(DeleteButtonPress event, Emitter<HomepageState> emit) async{
 
    emit(state.copyWith(status: HomepageStatus.loading));
    try {
      await deleteProductUsecase.execute(event.id);

      emit(state.copyWith(status: HomepageStatus.loaded));
      add(FetchProducts());
    } catch (e) {
      emit(state.copyWith(status: HomepageStatus.error));
    }
  }

  FutureOr<void> _addProduct(AddButtonPress event, Emitter<HomepageState> emit) async{
  
    emit(state.copyWith(status: HomepageStatus.loading));
    var product = ProductEntity(
        description: event.description,
        name: event.name,
        price: event.price,
        imageUrl: event.imageUrl);

    try {
      await addProductUsecase.execute(product);
      emit(state.copyWith(status: HomepageStatus.loaded));
      add(FetchProducts());
    } catch (e) {
      emit(state.copyWith(status: HomepageStatus.error));
    }
  }

  FutureOr<void> _updateProduct(UpdateButtonPress event, Emitter<HomepageState> emit) async{
   emit(state.copyWith(status: HomepageStatus.loading));
    var product = ProductEntity(
        id: event.id,
        description: event.description,
        name: event.name,
        price: event.price,
        imageUrl: event.imageUrl);

    try {
    await updateProductUsecase.execute(product);
      emit(state.copyWith(status: HomepageStatus.loaded));
      add(FetchProducts());
    } catch (e) {
      emit(state.copyWith(status: HomepageStatus.error));
    }

  }
}
