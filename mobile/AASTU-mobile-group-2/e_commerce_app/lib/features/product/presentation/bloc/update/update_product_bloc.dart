import 'dart:async';

import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/usecase/delete_product_usecase.dart';
import 'package:e_commerce_app/features/product/domain/usecase/update_product_usecase.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part 'update_product_event.dart';
part 'update_product_state.dart';

class UpdateProductBloc extends Bloc<ProductUpdatedEvent, UpdateProductState> {
  UpdateProduct updateProduct;
  DeleteProduct deleteProduct;
  UpdateProductBloc({required this.updateProduct,required this.deleteProduct})
      : super(UpdateProductInitial()) {
    on<ProductUpdated>(_onUpdated);
    on<UpdateInitiated>(_onUpdateInitiated);
    on<ProductDeleted>(_onProductDeleted);
  }

  FutureOr<void> _onUpdated(
      ProductUpdated event, Emitter<UpdateProductState> emit) async {
    emit(UpdateProductLoading());
    final result = await updateProduct.execute(event.product);
    result.fold((failure) => emit(UpdateProductFail()),
        (product) => emit(UpdateProductSuccess()));
  }

  FutureOr<void> _onUpdateInitiated(
      UpdateInitiated event, Emitter<UpdateProductState> emit) {
    emit(Update());
  }

  FutureOr<void> _onProductDeleted(event, Emitter<UpdateProductState> emit) async {
    emit(UpdateProductLoading());
    final result = await deleteProduct.execute(event.product);
    result.fold((failure) => emit(UpdateProductFail()),
        (product) => emit(UpdateProductSuccess()));
  }
  }



