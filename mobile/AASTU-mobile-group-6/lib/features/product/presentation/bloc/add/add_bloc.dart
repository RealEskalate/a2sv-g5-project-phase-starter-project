// add_bloc.dart


import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/add_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/add/add_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/add/add_state.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class addBloc extends Bloc<AddEvent, ProductState> {
  final AddProductUseCase addProductUseCase;

  addBloc(this.addProductUseCase) : super(ProductLoading()) {
    on<AddProductEvent>((event, emit) async {
      emit((ProductLoading()));
      var res = await addProductUseCase.call(event.product);
      print(res);
      print(res.fold((l) => l.message, (r) => r));
      res.fold((l)=> emit(ProductAddedFailure(error: l.message)), (r) => emit(ProductAddedSuccess(message: r)));

    });
    @override
  void onChange(Change<ProductState> change) {
    super.onChange(change);
    print('Current State: ${change.currentState}');
    print('Next State: ${change.nextState}');
  }
  }
}
