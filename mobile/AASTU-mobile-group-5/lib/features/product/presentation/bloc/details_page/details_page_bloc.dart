// details_page_bloc.dart
import 'package:bloc/bloc.dart';
import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../../core/failure/failure.dart';
import '../../../domain/entities/product.dart';
import '../../../domain/use_case/delete_product.dart';
import '../../../domain/use_case/get_product.dart';

part 'details_page_event.dart';
part 'details_page_state.dart';

class DetailsPageBloc extends Bloc<DetailsPageEvent, DetailsPageState> {
  final GetProduct getProduct;
  final DeleteProduct deleteProduct;

  DetailsPageBloc(this.getProduct, this.deleteProduct,) : super(DetailsPageInitialState()) {
    on<FetchProductByIdEvent>(_onFetchProductByIdEvent);
    on<DeleteDetailsEvent>(_onDeleteDetailsEvent);
  }

  Future<void> _onFetchProductByIdEvent(
      FetchProductByIdEvent event, Emitter<DetailsPageState> emit) async {
      emit(DetailsPageLoadingState());
      final product = await getProduct(event.params);
      product.fold((l)=>emit(DetailsPageErrorState(l.message)), (r) => emit(DetailsPageLoadedState(r)));
  }

  Future<void> _onDeleteDetailsEvent(
      DeleteDetailsEvent event, Emitter<DetailsPageState> emit) async {
    try {
      final Either<Failure, void> result = await deleteProduct(event.params);
      result.fold(
        (failure) => emit(DetailsPageErrorState(failure.toString())),
        (_) => emit(DetailsPageDeletedState()),
      );
    } catch (e) {
      emit(DetailsPageErrorState(e.toString()));
    }
  }
}
