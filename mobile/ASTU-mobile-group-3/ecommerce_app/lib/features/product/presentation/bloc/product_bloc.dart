import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/constants/constants.dart';
import '../../domain/entities/product.dart';
import '../../domain/usecases/delete_product_usecase.dart';
import '../../domain/usecases/get_all_products_usecase.dart';
import '../../domain/usecases/get_product_usecase.dart';
import '../../domain/usecases/insert_product_usecase.dart';
import '../../domain/usecases/update_product_usecase.dart';
import 'product_events.dart';
import 'product_states.dart';

class ProductBloc extends Bloc<ProductEvents, ProductStates> {
  final GetProductUseCase getProductUseCase;
  final GetAllProductUseCase getAllProductUseCase;
  final UpdateProductUsecase updateProductUsecase;
  final DeleteProductUseCase deleteProductUseCase;
  final InsertProductUseCase insertProductUseCase;
  ProductBloc({
    required this.getAllProductUseCase,
    required this.deleteProductUseCase,
    required this.getProductUseCase,
    required this.insertProductUseCase,
    required this.updateProductUsecase,
  }) : super(InitialState()) {
    on<GetSingleProductEvents>((event, emit) async {
      emit(LoadingState());
      final result = await getProductUseCase.execute(event.id);
      result.fold((failure) {
        emit(ErrorState(message: failure.message));
      }, (data) {
        emit(LoadedSingleProductState(productEntity: data));
      });
    });

    on<LoadAllProductEvents>((event, emit) async {
      emit(LoadingState());

      final result = await getAllProductUseCase.execute();

      result.fold((failure) {
        emit(ErrorState(message: failure.message));
      }, (data) {
        emit(LoadedAllProductState(data: data));
      });
    });

    on<InsertProductEvent>((event, emit) async {
      emit(LoadingState());
      ProductEntity entity = ProductEntity(
        id: '',
        name: event.name,
        description: event.description,
        price: int.parse(event.price),
        imageUrl: event.imageUrl.path,
      );
      final result = await insertProductUseCase.execute(entity);

      result.fold((failure) {
        emit(ErrorState(message: failure.message));
      }, (data) {
        emit(SuccessfullState(message: AppData.getMessage(data)));
      });
    });

    on<UpdateProductEvent>((event, emit) async {
      emit(LoadingState());

      ProductEntity entity = ProductEntity(
        id: event.id,
        name: event.name,
        description: event.description,
        price: int.parse(event.price),
        imageUrl: '',
      );
      final result = await updateProductUsecase.execute(entity);

      result.fold((failure) {
        emit(ErrorState(message: failure.message));
      }, (data) {
        emit(SuccessfullState(message: AppData.getMessage(data)));
      });
    });

    on<DeleteProductEvent>((event, emit) async {
      emit(LoadingState());
      final result = await deleteProductUseCase.execute(event.id);
      result.fold((failure) {
        emit(ErrorState(message: failure.message));
      }, (data) {
        emit(SuccessfullState(message: AppData.getMessage(data)));
      });
    });

    on<RefreshEvent>((event, emit) {
      if (state is LoadedAllProductState) {
        emit(LoadedAllProductState(data: [...state.data]));
      }
    });
  }
}
