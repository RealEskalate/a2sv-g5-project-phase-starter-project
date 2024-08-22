import 'dart:async';

import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/usecase/insert_product_usecase.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part 'insert_product_event.dart';
part 'insert_product_state.dart';

class InsertProductBloc extends Bloc<ProductInsertedEvent, InsertProductState> {
  InsertProduct insertProductUsecase;
  InsertProductBloc({required this.insertProductUsecase})
      : super(InsertProductInitial()) {
    on<ProductInserted>(_onInserted);
    on<InsertInitial>(_onInsertInitial);
  }

  FutureOr<void> _onInserted(
      ProductInserted event, Emitter<InsertProductState> emit) async {
    emit(InsertedProductLoading());
    final result = await insertProductUsecase.execute(event.product);
    result.fold((failure) => emit(InsertedProductFail()),
        (products) => emit(InsertedProductSuccess()));
         emit(InsertProductInitial());
  }

  FutureOr<void> _onInsertInitial(
      InsertInitial event, Emitter<InsertProductState> emit) {
    emit(InsertProductInitial());
  }
}
