import 'dart:async';
import 'package:bloc/bloc.dart';
// import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
// import '../../../../../core/failure/failure.dart';
import '../../../data/models/product_model.dart';
import '../../../domain/use_case/delete_product.dart';
import '../../../domain/use_case/update_product.dart';

part 'update_page_event.dart';
part 'update_page_state.dart';

class UpdatePageBloc extends Bloc<UpdatePageEvent, UpdatePageState> {
  final UpdateProduct updateProduct;
  final DeleteProduct deleteProduct;

  UpdatePageBloc(this.updateProduct, this.deleteProduct) : super(UpdatePageInitialState()) {
    on<UpdateProductEvent>(_onUpdateProduct);
    on<DeleteProductEvent>(_onDeleteProduct);
  }

  Future<void> _onUpdateProduct(UpdateProductEvent event, Emitter<UpdatePageState> emit) async {
  try {
    emit(UpdatePageSubmittingState());
    final result = await updateProduct(UpdateProductParams(event.product));

    result.fold(
      (failure) => emit(UpdatePageErrorState(failure.toString())),
      (product) => emit(UpdatePageSubmittedState(product)),
    );
  } catch (error) {
    emit(UpdatePageErrorState(error.toString()));
  }
}




  Future<void> _onDeleteProduct(DeleteProductEvent event, Emitter<UpdatePageState> emit) async {
    emit(UpdatePageSubmittingState());
    try {
      await deleteProduct(DeleteProductParams(event.productId));
      emit(UpdatePageDeletedState());
    } catch (e) {
      emit(UpdatePageErrorState(e.toString()));
    }
  }
}
