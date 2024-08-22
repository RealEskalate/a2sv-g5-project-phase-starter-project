// update_bloc.dart


import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/add_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/update_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/update/bloc/update_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/update/bloc/update_state.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class UpdateBloc extends Bloc<UpdateEvent, UpdateState> {
  final UpdateUsecase updateUsecase;

  UpdateBloc(this.updateUsecase) : super(UpdateLoading()) {
    on<UpdateProductEvent>((event, emit) async {
      var res = await updateUsecase.call(event.product);
      emit((UpdateLoading()));
      res.fold((l)=> emit(UpdateFailiure(error: l.message)), (r) => emit(UpdateSuccess()));

    });
  //   @override
  // void onChange(Change<UpdateState> change) {
  //   super.onChange(change);
  //   print('Current State: ${change.currentState}');
  //   print('Next State: ${change.nextState}');
  // }
  }
}
