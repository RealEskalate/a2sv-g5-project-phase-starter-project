// add_page_bloc.dart
import 'package:bloc/bloc.dart';
import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../../core/failure/failure.dart';
import '../../../data/models/product_model.dart';
import '../../../domain/use_case/add_product.dart';

part 'add_page_event.dart';
part 'add_page_state.dart';

class AddPageBloc extends Bloc<AddPageEvent, AddPageState> {
  final AddProduct addProduct;

  AddPageBloc(this.addProduct) : super(AddPageInitialState()) {
    on<AddProductEvent>(_onAddProductEvent);
  }

  Future<void> _onAddProductEvent(
      AddProductEvent event, Emitter<AddPageState> emit) async {
    try {
      emit(AddPageSubmittingState());
      final Either<Failure, ProductModel> result = await addProduct(
        AddProductParams(event.product, event.imagePath),
      );
      result.fold(
        (failure) => emit(AddPageErrorState(failure.toString())),
        (product) => emit(AddPageSubmittedState(product)),
      );
    } catch (e) {
      emit(AddPageErrorState(e.toString()));
    }
  }

  
  
}
